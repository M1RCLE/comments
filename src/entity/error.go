package entity

import (
	"fmt"
)

type ProcessError struct {
	Message string
}

func (e *ProcessError) Error() string {
	return fmt.Sprintf("Comment Error: %s", e.Message)
}
