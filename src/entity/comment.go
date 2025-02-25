package entity

type Comment struct {
	ID              int        `json:"id"`
	UserID          int        `json:"userId"`
	PostID          int        `json:"postId"`
	Body            string     `json:"body"`
	ParentId        *int       `json:"parent"`
	Indentation     int        `json:"indentation"`
	RelatedComments []*Comment `json:"relatedComments"`
}
