package service

import (
	"context"
	"errors"

	"github.com/M1RCLE/comments/src/entity"
	"github.com/M1RCLE/comments/src/repository/contract"
)

type PostService struct {
	repository contract.Repository
}

func NewPostService(repo contract.Repository) *PostService {
	return &PostService{
		repository: repo,
	}
}

func (ps *PostService) CreatePost(ctx context.Context, post entity.Post) (*entity.Post, error) {
	if post.Body == "" {
		return nil, errors.New("post body cannot be empty")
	}
	if post.UserId == 0 {
		return nil, errors.New("invalid user ID")
	}

	return ps.repository.CreatePost(ctx, post)
}

func (ps *PostService) GetPosts(ctx context.Context, limit *int, offset *int) ([]*entity.Post, error) {
	pagination := entity.Pagination{
		Limit:  getOrDefault(limit, 10),
		Offset: getOrDefault(offset, 0),
	}
	return ps.repository.GetPosts(ctx, pagination)
}

func (ps *PostService) GetPostById(ctx context.Context, postID int) (*entity.Post, error) {
	return ps.repository.GetPostById(ctx, postID)
}

func (ps *PostService) SwitchPostAllowance(ctx context.Context, postID int) error {
	return ps.repository.SwitchPostAllowance(ctx, postID)
}

func (ps *PostService) DeletePost(ctx context.Context, postID int) error {
	return ps.repository.DeletePost(ctx, postID)
}

func getOrDefault(val *int, defaultValue int) int {
	if val != nil {
		return *val
	}
	return defaultValue
}
