package service

import (
	"context"

	"github.com/M1RCLE/comments/src/entity"
	contract "github.com/M1RCLE/comments/src/repository"
)

type CommentService struct {
	repository contract.Repository
}

func (cs *CommentService) CreateComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error) {
	panic("not implemented")
}

func (cs *CommentService) CreateSubComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error) {
	panic("not implemented")
}

func (cs *CommentService) GetCommentById(ctx context.Context, commentID string) (*entity.Comment, error) {
	panic("not implemented")
}

func (cs *CommentService) GetComments(ctx context.Context, limit *int) ([]*entity.Comment, error) {
	panic("not implemented")
}

func (cs *CommentService) DeleteComment(ctx context.Context, commentID string) error {
	panic("not implemented")
}
