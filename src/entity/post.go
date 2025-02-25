package entity

type Post struct {
	ID              int        `json:"id"`
	Body            string     `json:"body"`
	UserID          int        `json:"userId"`
	Comments        []*Comment `json:"comments"`
	CommentsAllowed bool       `json:"commentsAllowed"`
}
