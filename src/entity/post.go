package entity

import "time"

type Post struct {
	Id              int        `json:"id"`
	Body            string     `json:"body"`
	UserId          int        `json:"userId"`
	Comments        []*Comment `json:"comments"`
	CommentsAllowed bool       `json:"commentsAllowed"`
	CreatedAt       time.Time  `json:"createdAt"`
}
