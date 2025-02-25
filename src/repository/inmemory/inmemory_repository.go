package repository

import (
	"context"
	"errors"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/M1RCLE/comments/src/entity"
	"github.com/M1RCLE/comments/src/repository/data"
)

type storage map[int]*data.Post

type InmemoryRepository struct {
	sync.RWMutex
	idUpscaler int
	storage    storage
}

func NewStorage() *InmemoryRepository {
	return &InmemoryRepository{
		idUpscaler: 1,
		storage:    make(storage),
	}
}

func (r *InmemoryRepository) CreatePost(_ context.Context, post entity.Post) (*entity.Post, error) {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.storage[post.Id]; exists {
		return nil, &entity.ProcessError{Message: "Post with ID:" + strconv.Itoa(post.Id) + " already exists"}
	}

	newPost := data.NewPost(r.idUpscaler, post.Body, post.UserId, post.CommentsAllowed)
	r.idUpscaler++
	r.storage[newPost.Id] = newPost

	return PostFromData(newPost), nil
}

func (r *InmemoryRepository) GetPosts(_ context.Context, pagination entity.Pagination) ([]*entity.Post, error) {
	r.RLock()
	defer r.RUnlock()

	posts := make([]*entity.Post, 0, len(r.storage))

	for _, post := range r.storage {
		posts = append(posts, PostFromData(post))
	}

	start := pagination.Offset
	if start > len(posts) {
		return []*entity.Post{}, nil
	}

	return PartialSort(posts, pagination), nil
}

func (r *InmemoryRepository) GetPostById(_ context.Context, postId int) (*entity.Post, error) {
	r.RLock()
	defer r.RUnlock()

	post, exists := r.storage[postId]
	if !exists {
		return nil, &entity.ProcessError{Message: "post not found"}
	}

	return PostFromData(post), nil
}

func (r *InmemoryRepository) DeletePost(_ context.Context, postId int) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.storage[postId]; !exists {
		return errors.New("post not found")
	}

	for value := range r.storage[postId].Comments {
		r.storage[postId].DeleteComment(value)
	}
	delete(r.storage, postId)
	return nil
}

func (r *InmemoryRepository) SwitchPostAllowance(_ context.Context, postId int) error {
	r.Lock()
	defer r.Unlock()

	post, exists := r.storage[postId]
	if !exists {
		return &entity.ProcessError{Message: "post not found"}
	}

	post.CommentsAllowed = !post.CommentsAllowed
	return nil
}

func (r *InmemoryRepository) CreateComment(_ context.Context, comment entity.Comment) (*entity.Comment, error) {
	r.Lock()
	defer r.Unlock()

	post, exists := r.storage[comment.PostId]
	if !exists {
		return nil, &entity.ProcessError{Message: "post not found"}
	}

	if !post.CommentsAllowed {
		return nil, &entity.ProcessError{Message: "comments are disabled for this post"}
	}

	return post.AddComment(&comment), nil
}

func (r *InmemoryRepository) CreateSubComment(_ context.Context, comment entity.Comment) (*entity.Comment, error) {
	r.Lock()
	defer r.Unlock()

	post, exists := r.storage[comment.PostId]
	if !exists {
		return nil, &entity.ProcessError{Message: "post not found"}
	}

	if _, err := post.Comments[*comment.ParentId]; !err {
		return nil, &entity.ProcessError{Message: "parent comment not found"}
	}

	return post.AddSubComment(&comment), nil
}

func (r *InmemoryRepository) GetComments(_ context.Context, pagination entity.Pagination) ([]*entity.Comment, error) {
	r.RLock()
	defer r.RUnlock()

	comments := []*entity.Comment{}

	for _, post := range r.storage {
		for _, comment := range post.Comments {
			comments = append(comments, comment)
		}
	}

	start := pagination.Offset
	if start > len(comments) {
		return []*entity.Comment{}, nil
	}

	return PartialSort(comments, pagination), nil
}

func (r *InmemoryRepository) DeleteComment(_ context.Context, commentId int) error {
	r.Lock()
	defer r.Unlock()

	for _, post := range r.storage {
		if _, exists := post.Comments[commentId]; exists {
			delete(post.Comments, commentId)
			return nil
		}
	}

	return &entity.ProcessError{Message: "comment not found"}
}

func PostFromData(post *data.Post) *entity.Post {
	comments := make([]*entity.Comment, 0, len(post.Comments))
	for _, comment := range post.Comments {
		comments = append(comments, comment)
	}

	return &entity.Post{
		Id:              post.Id,
		Body:            post.Body,
		UserId:          post.UserId,
		Comments:        comments,
		CommentsAllowed: post.CommentsAllowed,
	}
}

type Creatable interface {
	GetCreationTime() time.Time
}

func PartialSort[E Creatable](slice []E, pagination entity.Pagination) []E {
	n := len(slice)
	offset := pagination.Offset
	limit := pagination.Limit

	if offset < 0 || offset >= n {
		return slice
	}

	end := offset + limit
	if end > n {
		end = n
	}

	sortedPart := append([]E{}, slice[offset:end]...)
	sort.Slice(sortedPart, func(i, j int) bool {
		return sortedPart[i].GetCreationTime().Before(sortedPart[j].GetCreationTime())
	})

	result := append([]E{}, slice...)
	copy(result[offset:end], sortedPart)

	return result
}
