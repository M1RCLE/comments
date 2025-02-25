package interf

import (
	"context"

	"github.com/M1RCLE/comments/src/entity"
)

type Post interface {
	CreatePost(ctx context.Context, post entity.Post) (*entity.Post, error)
	GetPosts(ctx context.Context, limit *int, offset *int) ([]*entity.Post, error)
	GetPostById(ctx context.Context, postID int) (*entity.Post, error)
	SwitchPostAllowance(ctx context.Context, postID int) error
	DeletePost(ctx context.Context, postID int) error
}
