package interf

import (
	"context"

	"github.com/M1RCLE/comments/src/entity"
)

type Comment interface {
	CreateComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error)
	CreateSubComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error)
	GetCommentById(ctx context.Context, commentID int) (*entity.Comment, error)
	GetAllComments(ctx context.Context, commentID int) (*entity.Comment, error)
	GetComments(ctx context.Context, limit *int, offset *int) ([]*entity.Comment, error)
	DeleteComment(ctx context.Context, commentID int) error
}
