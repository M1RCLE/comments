package entity

import "time"

type Comment struct {
	ID              int        `json:"id" db:"id"`
	UserId          int        `json:"userId" db:"user_id"`
	PostId          int        `json:"postId" db:"post_id"`
	Body            string     `json:"body" db:"body"`
	ParentId        *int       `json:"parent" db:"parent"`
	CreationTime    time.Time  `json:"creationTime" db:"creation_time"`
	RelatedComments []*Comment `json:"relatedComments" db:"related_comments"`
}

func (c *Comment) GetCreationTime() time.Time {
	return c.CreationTime
}
