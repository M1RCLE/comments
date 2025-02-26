package service

import (
	"context"
	"errors"
	"sync"

	"github.com/M1RCLE/comments/src/entity"
	"github.com/M1RCLE/comments/src/repository/contract"
	interf "github.com/M1RCLE/comments/src/service/interface"
)

type SubscriptionService struct {
	repository    contract.Repository
	subscriptions map[int]map[int]chan *entity.Comment // postId -> userId -> channel
	mu            sync.RWMutex
}

func NewSubscriptionService(repo contract.Repository) interf.Subscription {
	return &SubscriptionService{
		repository:    repo,
		subscriptions: make(map[int]map[int]chan *entity.Comment),
	}
}

func (s *SubscriptionService) RegisterSubscription(ctx context.Context, userId int, postId int) (<-chan *entity.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.repository.GetPostById(ctx, postId)
	if err != nil {
		return nil, errors.New("post not found")
	}

	if s.subscriptions[postId] == nil {
		s.subscriptions[postId] = make(map[int]chan *entity.Comment)
	}

	commentChannel := make(chan *entity.Comment, 10)
	s.subscriptions[postId][userId] = commentChannel

	return commentChannel, nil
}

func (s *SubscriptionService) UnregisterSubscription(ctx context.Context, userId int, postId int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.subscriptions[postId]; !exists {
		return errors.New("no subscriptions for this post")
	}

	if ch, exists := s.subscriptions[postId][userId]; exists {
		close(ch)
		delete(s.subscriptions[postId], userId)

		if len(s.subscriptions[postId]) == 0 {
			delete(s.subscriptions, postId)
		}
		return nil
	}

	return errors.New("subscription not found")
}

func (s *SubscriptionService) NotifySubscribers(ctx context.Context, postId int, comment *entity.Comment) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if subs, exists := s.subscriptions[postId]; exists {
		for _, ch := range subs {
			select {
			case ch <- comment:
			default:
			}
		}
	}

	return nil
}
