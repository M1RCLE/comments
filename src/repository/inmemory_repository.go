package repository

import (
	"context"
	"strconv"
	"sync"

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
	newPost := data.NewPost(r.idUpscaler, post.Body, post.UserId, post.CommentsAllowed)
	r.idUpscaler++
	if _, find := r.storage[post.Id]; find {
		return nil, &entity.ProcessError{Message: "Post with ID:" + strconv.Itoa(post.Id) + " already exists"}
	}
	r.storage[post.Id] = newPost
	return PostFromData(newPost), nil
}
func (r *InmemoryRepository) GetPosts(ctx context.Context, limit *int) ([]*entity.Post, error) {
	r.RLock()
	defer r.RUnlock()

}

func (r *InmemoryRepository) GetPostById(ctx context.Context, postID string) (*entity.Post, error) {
	panic("not impleamented")
}
func (r *InmemoryRepository) DeletePost(ctx context.Context, postID string) error {
	panic("not impleamented")
}

func (r *InmemoryRepository) CreateComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error) {
	panic("not impleamented")
}
func (r *InmemoryRepository) CreateSubComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error) {
	panic("not impleamented")
}
func (r *InmemoryRepository) GetComments(ctx context.Context, limit *int) ([]*entity.Comment, error) {
	panic("not impleamented")
}
func (r *InmemoryRepository) GetCommentById(ctx context.Context, commentID string) (*entity.Comment, error) {
	panic("not impleamented")
}
func (r *InmemoryRepository) DeleteComment(ctx context.Context, commentID string) error {
	panic("not impleamented")
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
