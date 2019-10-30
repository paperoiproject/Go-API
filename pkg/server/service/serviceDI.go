package service

import (
	"prote-API/pkg/server/repository"
	"prote-API/pkg/server/service/scenario"
	"prote-API/pkg/server/service/scene"
	"prote-API/pkg/server/service/scenes"
	"prote-API/pkg/server/service/test"
)

// service サービスの構造体
type service struct {
	TestService     *test.TestService
	SceneService    *scene.SceneService
	ScenarioService *scenario.ScenarioService
	ScenesService   *scenes.ScenesService
}

// Service サービスの生成(依存関係の解決)
var Service = service{
	TestService:     &test.TestService{SceneTable: &repository.Scene{}},
	SceneService:    &scene.SceneService{SceneTable: &repository.Scene{}},
	ScenesService:   &scenes.ScenesService{ScenesTable: &repository.Scenes{}},
	ScenarioService: &scenario.ScenarioService{ScenarioTable: &repository.ScenarioTable{}},
}
