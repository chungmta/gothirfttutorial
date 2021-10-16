package store

import (
	"context"
	"thirfttutorial/src/modal"
)

type Store interface {
	CreateStudent(ctx context.Context, s modal.Student) error
	CreateClass(ctx context.Context, s modal.Class) error
	GetStudentsByClassID(ctx context.Context, class modal.Class) ([]modal.Student, error)
}
