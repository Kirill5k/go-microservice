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

type IdMissmatchError struct {
	PathID uuid.UUID
	BodyID uuid.UUID
}

func (e *IdMissmatchError) Error() string {
	return fmt.Sprintf("id in path %s is different from id in body %s", e.PathID, e.BodyID)
}
