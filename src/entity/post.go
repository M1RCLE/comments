package entity

import "time"

type Post struct {
	Id              int        `json:"id" db:"id"`
	Body            string     `json:"body" db:"body"`
	UserId          int        `json:"userId" db:"user_id"`
	Comments        []*Comment `json:"comments" db:"comments"`
	CommentsAllowed bool       `json:"commentsAllowed" db:"comments_allowed"`
	CreationTime    time.Time  `json:"createdAt" db:"creation_time"`
}

func (p *Post) GetCreationTime() time.Time {
	return p.CreationTime
}
