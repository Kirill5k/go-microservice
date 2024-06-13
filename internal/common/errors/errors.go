package errors

import (
	"errors"
	_ "errors"
	"fmt"
	"github.com/google/uuid"
)

func Is(err, target error) bool {
	return errors.Is(err, target)
}

type AppError struct {
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

type ConflictError struct {
	Detail string
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("attempted to create a record with an existing key: %s", e.Detail)
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
