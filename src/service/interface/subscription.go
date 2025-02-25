package interf

import (
	"context"

	"github.com/M1RCLE/comments/src/entity"
)

type Subscription interface {
	RegisterSubscription(ctx context.Context, userId int, postId int) (<-chan *entity.Comment, error)
}
