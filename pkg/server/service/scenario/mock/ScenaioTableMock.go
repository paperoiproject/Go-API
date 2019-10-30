package mock

import (
	"prote-API/pkg/server/repository"
	"time"
)

// ScenarioTable ScenarioTableのmock
type ScenarioTable interface {
	SelectAll() ([]repository.ScenarioRow, error)
	Insert(name string, date time.Time) error
	Make(name string, date time.Time, texts []string, acts []string, images []string) error
}
