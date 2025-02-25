package data

import (
	"sync"

	"github.com/M1RCLE/comments/src/entity"
)

// Коментарии хранятся в виде мапы, где ключ - это id комментария, а значение - сам комментарий
type Post struct {
	sync.Mutex
	Id              int                     `json:"id"`
	Body            string                  `json:"body"`
	UserId          int                     `json:"userId"`
	Comments        map[int]*entity.Comment `json:"comments"`
	CommentsAllowed bool                    `json:"commentsAllowed"`
	IdUpscaler      int
}

func NewPost(id int, body string, userId int, commentsAllowed bool) *Post {
	return &Post{
		Id:              id,
		Body:            body,
		UserId:          userId,
		Comments:        make(map[int]*entity.Comment),
		CommentsAllowed: commentsAllowed,
		IdUpscaler:      1,
	}
}

func (post *Post) AddComment(comment *entity.Comment) *entity.Comment {
	post.Lock()
	defer post.Unlock()

	if _, ok := post.Comments[comment.ID]; ok {
		return nil
	}

	comment.ID = post.IdUpscaler
	post.IdUpscaler++
	post.Comments[comment.ID] = comment

	return comment
}

func (post *Post) AddSubComment(comment *entity.Comment) *entity.Comment {
	post.Lock()
	defer post.Unlock()

	if _, ok := post.Comments[*comment.ParentId]; !ok {
		return nil
	}

	if _, ok := post.Comments[*comment.ParentId]; !ok {
		return nil
	}

	comment.ID = post.IdUpscaler
	post.IdUpscaler++
	post.Comments[comment.ID] = comment

	return comment
}

func (post *Post) DeleteComment(commentId int) []*entity.Comment {
	post.Lock()
	defer post.Unlock()

	comment, ok := post.Comments[commentId]
	if !ok {
		return nil
	}

	for _, repComment := range comment.RelatedComments {
		post.Unlock()
		post.DeleteComment(repComment.ID)
		post.Lock()
	}
	delete(post.Comments, commentId)

	return nil
}
