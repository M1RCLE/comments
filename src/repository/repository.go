package contract

import (
	"context"

	"github.com/M1RCLE/comments/src/entity"
)

type Repository interface {
	CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error)
	GetPosts(ctx context.Context, limit *int) ([]*entity.Post, error)
	GetPostById(ctx context.Context, postID string) (*entity.Post, error)
	DeletePost(ctx context.Context, postID string) error

	CreateComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, error)
	CreateSubComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, error)
	GetComments(ctx context.Context, limit *int) ([]*entity.Comment, error)
	GetCommentById(ctx context.Context, commentID string) (*entity.Comment, error)
	DeleteComment(ctx context.Context, commentID string) error
}
