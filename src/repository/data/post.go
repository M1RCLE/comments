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

func (post *Post) AddComment(commentId int, comment *entity.Comment) {
	post.Comments[commentId] = comment
}
