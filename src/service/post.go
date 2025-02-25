package service

import (
	"context"

	"github.com/M1RCLE/comments/src/entity"
	contract "github.com/M1RCLE/comments/src/repository"
)

type PostService struct {
	repository contract.Repository
}

func (ps *PostService) CreatePost(ctx context.Context, post entity.Post) (*entity.Post, error) {
	panic("not implemented")
}

func (ps *PostService) GetPosts(ctx context.Context, limit int) ([]*entity.Post, error) {
	panic("not implemented")
}
