// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type CommentInput struct {
	UserID       int       `json:"userId"`
	PostID       int       `json:"postId"`
	CreationTime time.Time `json:"creationTime"`
	Body         string    `json:"body"`
}

type Mutation struct {
}

type PostInput struct {
	UserID          int    `json:"userId"`
	Body            string `json:"body"`
	CommentsAllowed bool   `json:"commentsAllowed"`
}

type Query struct {
}

type SubCommentInput struct {
	UserID       int       `json:"userId"`
	PostID       int       `json:"postId"`
	ParentID     int       `json:"parentId"`
	CreationTime time.Time `json:"creationTime"`
	Body         string    `json:"body"`
}

type Subscription struct {
}
