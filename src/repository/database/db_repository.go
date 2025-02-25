package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/M1RCLE/comments/src/entity"
	"github.com/Masterminds/squirrel"
)

type DatabaseRepository struct {
	postStorage    *Storage[entity.Post]
	commentStorage *Storage[entity.Comment]
	sqlBuilder     squirrel.StatementBuilderType
}

func NewDatabaseRepository(postStorage *Storage[entity.Post], commentStorage *Storage[entity.Comment]) *DatabaseRepository {
	return &DatabaseRepository{
		postStorage:    postStorage,
		commentStorage: commentStorage,
		sqlBuilder:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

// Create a new post
func (d *DatabaseRepository) CreatePost(ctx context.Context, post entity.Post) (*entity.Post, error) {
	query, args, err := d.sqlBuilder.Insert("post").
		Columns("user_id", "body", "comments_allowed", "creation_time").
		Values(post.UserId, post.Body, post.CommentsAllowed, post.CreationTime).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}

	createdPost, err := d.postStorage.Insert(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error inserting post: %w", err)
	}

	return &createdPost, nil
}

// Get multiple posts with pagination
func (d *DatabaseRepository) GetPosts(ctx context.Context, pagination entity.Pagination) ([]*entity.Post, error) {
	var posts []*entity.Post

	queryBuilder := d.sqlBuilder.Select("id", "user_id", "body", "comments_allowed", "creation_time").
		From("post").
		OrderBy("timestamp DESC")

	queryBuilder = queryBuilder.Limit(uint64(pagination.Limit))
	queryBuilder = queryBuilder.Offset(uint64(pagination.Offset))

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}

	rows, err := d.postStorage.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error fetching posts: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		post := &entity.Post{}
		if err := rows.StructScan(post); err != nil {
			return nil, fmt.Errorf("error scanning post: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// Get a single post by ID
func (d *DatabaseRepository) GetPostById(ctx context.Context, postId int) (*entity.Post, error) {
	query, args, err := d.sqlBuilder.Select("id", "user_id", "content", "is_open", "timestamp").
		From("post").
		Where(squirrel.Eq{"id": postId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}

	post, err := d.postStorage.Insert(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("post not found: %w", err)
		}
		return nil, fmt.Errorf("error fetching post by ID: %w", err)
	}

	return &post, nil
}

// Delete a post by ID
func (d *DatabaseRepository) DeletePost(ctx context.Context, postId int) error {
	query, args, err := d.sqlBuilder.Delete("post").
		Where(squirrel.Eq{"id": postId}).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building SQL query: %w", err)
	}

	return d.postStorage.Delete(ctx, query, args...)
}

// Switch the comment allowance state of a post
func (d *DatabaseRepository) SwitchPostAllowance(ctx context.Context, postId int) error {
	query, args, err := d.sqlBuilder.Update("post").
		Set("is_open", squirrel.Expr("NOT is_open")). // Toggle the value
		Where(squirrel.Eq{"id": postId}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = d.postStorage.ExecuteUpdate(ctx, query, args...)
	return err
}

// Create a new comment
func (d *DatabaseRepository) CreateComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error) {
	query, args, err := d.sqlBuilder.Insert("comment").
		Columns("post_id", "user_id", "body", "creation_time").
		Values(comment.PostId, comment.UserId, comment.Body, comment.CreationTime).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}

	createdComment, err := d.commentStorage.Insert(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error inserting comment: %w", err)
	}

	return &createdComment, nil
}

// Create a sub-comment (reply)
func (d *DatabaseRepository) CreateSubComment(ctx context.Context, comment entity.Comment) (*entity.Comment, error) {
	if comment.ParentId == nil {
		return nil, errors.New("parent ID is required for sub-comments")
	}

	query, args, err := d.sqlBuilder.Insert("comment").
		Columns("parent_id", "post_id", "user_id", "body", "creation_time").
		Values(*comment.ParentId, comment.PostId, comment.UserId, comment.Body, comment.CreationTime).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}

	createdComment, err := d.commentStorage.Insert(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error inserting sub-comment: %w", err)
	}

	return &createdComment, nil
}

// Get multiple comments with pagination
func (d *DatabaseRepository) GetComments(ctx context.Context, pagination entity.Pagination) ([]*entity.Comment, error) {
	var comments []*entity.Comment

	queryBuilder := d.sqlBuilder.Select("id", "parent_id", "post_id", "user_id", "content", "timestamp").
		From("comment").
		OrderBy("timestamp")

	queryBuilder = queryBuilder.Limit(uint64(pagination.Limit))
	queryBuilder = queryBuilder.Offset(uint64(pagination.Offset))

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}

	rows, err := d.commentStorage.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error fetching comments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		comment := &entity.Comment{}
		if err := rows.StructScan(comment); err != nil {
			return nil, fmt.Errorf("error scanning comment: %w", err)
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

// Get a single comment by ID
func (d *DatabaseRepository) GetCommentById(ctx context.Context, commentId int) (*entity.Comment, error) {
	query, args, err := d.sqlBuilder.Select("id", "parent_id", "post_id", "user_id", "content", "timestamp").
		From("comment").
		Where(squirrel.Eq{"id": commentId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}

	comment, err := d.commentStorage.Insert(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("comment not found: %w", err)
		}
		return nil, fmt.Errorf("error fetching comment by ID: %w", err)
	}

	return &comment, nil
}

// Delete a comment by ID
func (d *DatabaseRepository) DeleteComment(ctx context.Context, commentId int) error {
	query, args, err := d.sqlBuilder.Delete("comment").
		Where(squirrel.Eq{"id": commentId}).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building SQL query: %w", err)
	}

	return d.commentStorage.Delete(ctx, query, args...)
}
