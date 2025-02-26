package contract

import (
	"context"

	"github.com/M1RCLE/comments/src/entity"
)

type Repository interface {
	CreatePost(ctx context.Context, post entity.Post) (*entity.Post, error)
	GetPosts(ctx context.Context, pagination entity.Pagination) ([]*entity.Post, error)
	GetPostById(ctx context.Context, postId int) (*entity.Post, error)
	DeletePost(ctx context.Context, postId int) error

	SwitchPostAllowance(ctx context.Context, postId int) error
	CreateComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error)
	CreateSubComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error)
	GetComments(ctx context.Context, pagination entity.Pagination) ([]*entity.Comment, error)
	GetCommentById(ctx context.Context, commentId int) (*entity.Comment, error)
	DeleteComment(ctx context.Context, commentId int) error
}
