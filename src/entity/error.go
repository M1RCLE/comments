package entity

import (
	"fmt"
)

type CommentError struct {
	Message string
}

func (e *CommentError) Error() string {
	return fmt.Sprintf("Comment Error: %s", e.Message)
}
