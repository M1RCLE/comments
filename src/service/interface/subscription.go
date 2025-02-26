package interf

import (
	"context"

	"github.com/M1RCLE/comments/src/entity"
)

type Subscription interface {
	RegisterSubscription(ctx context.Context, userId int, postId int) (<-chan *entity.Comment, error)
	UnregisterSubscription(ctx context.Context, userId int, postId int) error
	NotifySubscribers(ctx context.Context, postId int, comment *entity.Comment) error
}
