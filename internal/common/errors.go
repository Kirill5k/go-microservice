package common

import (
	"fmt"
	"github.com/google/uuid"
)

type ConflictError struct {
}

func (e *ConflictError) Error() string {
	return "attempted to create a record with an existing key"
}

type NotFoundError struct {
	Entity string
	ID     uuid.UUID
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("unable to find %s with id %s", e.Entity, e.ID)
}

type InvalidIdError struct {
	ID string
}

func (e *InvalidIdError) Error() string {
	return fmt.Sprintf("%s is not a valid uuid", e.ID)
}
