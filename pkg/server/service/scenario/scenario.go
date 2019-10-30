package scenario

import (
	"prote-API/pkg/server/repository"
	"prote-API/pkg/server/service/scenario/mock"
	"sort"
	"time"
)

// ScenarioService /scenario以下のサービス
type ScenarioService struct {
	ScenarioTable mock.ScenarioTable
}

// List /scenario/listのサービス
func (scenarioService *ScenarioService) List() ([]repository.ScenarioRow, error) {
	scenarios, err := scenarioService.ScenarioTable.SelectAll()
	if err != nil {
		return nil, err
	}
	sort.Slice(scenarios, func(i, j int) bool {
		return scenarios[j].Date.Before(scenarios[i].Date)
	})
	return scenarios, err
}

// Add /scenario/addのサービス
func (scenarioService *ScenarioService) Add(name string, date time.Time) error {
	err := scenarioService.ScenarioTable.Insert(name, date)
	return err
}

// Make /scenario/makeのサービス
func (scenarioService *ScenarioService) Make(name string, date time.Time, acts []string, texts []string, images []string) error {
	err := scenarioService.ScenarioTable.Make(name, date, acts, texts, images)
	return err
}
