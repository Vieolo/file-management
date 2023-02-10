package filemanagement

import "fmt"

type NotFoundError struct {
	Target string
}

func (m *NotFoundError) Error() string {
	return fmt.Sprintf("%v was not found", m.Target)
}
