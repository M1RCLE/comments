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
	if len(comment.Body) > 2000 {
		return nil, &entity.CommentError{Message: "Comment is too long"}
	}
	post, err := cs.repository.GetPostById(ctx, comment.PostID)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, &entity.CommentError{Message: "Post not found"}
	}
	return cs.repository.CreateComment(ctx, &comment)
}

func (cs *CommentService) CreateSubComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error) {
	if len(comment.Body) > 2000 {
		return nil, &entity.CommentError{Message: "Comment is too long"}
	}
	commentPtr, err := cs.repository.GetCommentById(ctx, *comment.ParentId)
	if err != nil {
		return nil, err
	}
	if commentPtr == nil {
		return nil, &entity.CommentError{Message: "Comment not found"}
	}
	return cs.repository.CreateSubComment(ctx, &comment)
}

func (cs *CommentService) GetCommentById(ctx context.Context, commentID string) (*entity.Comment, error) {
	return cs.repository.GetCommentById(ctx, commentID)
}

func (cs *CommentService) GetComments(ctx context.Context, limit *int) ([]*entity.Comment, error) {
	return cs.repository.GetComments(ctx, limit)
}

func (cs *CommentService) DeleteComment(ctx context.Context, commentID string) error {
	_, err := cs.repository.GetCommentById(ctx, commentID)
	if err != nil {
		return err
	}
	return cs.repository.DeleteComment(ctx, commentID)
}
