package entity

import "time"

type Comment struct {
	ID              int        `json:"id"`
	UserId          int        `json:"userId"`
	PostId          int        `json:"postId"`
	Body            string     `json:"body"`
	ParentId        *int       `json:"parent"`
	Indentation     int        `json:"indentation"`
	RelatedComments []*Comment `json:"relatedComments"`
	CreatedAt       time.Time  `json:"createdAt"`
}
