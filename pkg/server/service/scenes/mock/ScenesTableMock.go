package mock

import (
	"prote-API/pkg/server/repository"
)

// ScenesTable SceneTableのmock
type ScenesTable interface {
	SelectAllRows() ([]repository.ScenesRow, error)
}
