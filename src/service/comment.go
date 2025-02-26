package service

import (
	"context"
	"errors"

	"github.com/M1RCLE/comments/src/entity"
	"github.com/M1RCLE/comments/src/repository/contract"
)

type CommentService struct {
	repository contract.Repository
}

func (cs *CommentService) CreateComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error) {
	if comment.Body == "" || len(comment.Body) > 2000 {
		return nil, &entity.ProcessError{Message: "comment body is required"}
	}
	if comment.PostId == 0 {
		return nil, &entity.ProcessError{Message: "post id is required"}
	}

	post, err := cs.repository.GetPostById(ctx, comment.PostId)
	if err != nil {
		return nil, err
	}

	if !post.CommentsAllowed {
		return nil, &entity.ProcessError{Message: "comments are disabled for this post"}
	}

	return cs.repository.CreateComment(ctx, comment)
}

func (cs *CommentService) CreateSubComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error) {
	if comment.Body == "" || len(comment.Body) > 2000 {
		return nil, errors.New("comment body cannot be empty")
	}
	if comment.PostId == 0 || (comment.ParentId != nil && *comment.ParentId == 0) {
		return nil, errors.New("invalid post ID or parent comment ID")
	}

	post, err := cs.repository.GetPostById(ctx, comment.PostId)
	if err != nil {
		return nil, err
	}
	if !post.CommentsAllowed {
		return nil, errors.New("comments are disabled for this post")
	}

	return cs.repository.CreateSubComment(ctx, comment)
}

func (cs *CommentService) GetAllComments(ctx context.Context, commentID int) (*entity.Comment, error) {
	panic("not implemented")
}

func (cs *CommentService) GetComments(ctx context.Context, limit *int, offset *int) ([]*entity.Comment, error) {
	panic("not implemented")
}

func (cs *CommentService) DeleteComment(ctx context.Context, commentID string) error {
	panic("not implemented")
}
