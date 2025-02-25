package interf

import (
	"context"

	"github.com/M1RCLE/comments/src/entity"
)

type Post interface {
	CreatePost(ctx context.Context, post entity.Post) (*entity.Post, error)
	GetPosts(ctx context.Context, limit *int) ([]*entity.Post, error)
	GetPostById(ctx context.Context, postID int) (*entity.Post, error)
	ChangePostCommentMod(ctx context.Context, postID string)
	DeletePost(ctx context.Context, postID int) error
}
