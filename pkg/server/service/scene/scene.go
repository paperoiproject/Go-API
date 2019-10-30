package scene

import (
	"prote-API/pkg/server/repository"
	"prote-API/pkg/server/service/scene/mock"
	"sort"
)

// SceneService /scene以下のサービス
type SceneService struct {
	SceneTable mock.SceneTable
}

// List /scene/listのサービス
func (sceneService *SceneService) List(name string) ([]repository.SceneRow, error) {
	scenes, err := sceneService.SceneTable.SelectRowsByName(name)
	if err != nil {
		return nil, err
	}
	sort.Slice(scenes, func(i, j int) bool {
		return scenes[i].Num < scenes[j].Num
	})
	return scenes, err
}

// Add /scene/addのサービス
func (sceneService *SceneService) Add(name string, num int, works []string, texts []string) error {
	err := sceneService.SceneTable.BulkInsert(name, num, works, texts)
	return err
}

// Delete /scene/deleteのサービス
func (sceneService *SceneService) Delete(name string) error {
	err := sceneService.SceneTable.Delete(name)
	return err
}

// Update /scene/deleteのサービス
func (sceneService *SceneService) Update(name string, num int, works []string, texts []string) error {
	err := sceneService.SceneTable.Delete(name)
	if err != nil {
		return err
	}
	err = sceneService.SceneTable.BulkInsert(name, num, works, texts)
	return err
}
