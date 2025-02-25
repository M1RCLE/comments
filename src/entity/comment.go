package entity

import "time"

type Comment struct {
	ID              int        `json:"id"`
	UserId          int        `json:"userId"`
	PostId          int        `json:"postId"`
	Body            string     `json:"body"`
	ParentId        *int       `json:"parent"`
	CreationTime    time.Time  `json:"creationTime"`
	RelatedComments []*Comment `json:"relatedComments"`
}

func (c *Comment) GetCreationTime() time.Time {
	return c.CreationTime
}
