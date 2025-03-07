package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	service "github.com/M1RCLE/comments/src/service/interface"
)

type Resolver struct {
	PostService         service.Post
	CommentService      service.Comment
	SubscriptionService service.Subscription
}

func NewResolver(postService service.Post, commentService service.Comment, subscriptionService service.Subscription) *Resolver {
	return &Resolver{
		PostService:         postService,
		CommentService:      commentService,
		SubscriptionService: subscriptionService,
	}
}
