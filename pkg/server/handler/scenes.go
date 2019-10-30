package handler

import (
	"log"
	"net/http"
	"prote-API/pkg/server/handler/response"
	"prote-API/pkg/server/service"
)

// GetHandleScenesList /scenes/listのハンドラ(シーンリストの変更)
func GetHandleScenesList(writer http.ResponseWriter, request *http.Request) {
	scenes, err := service.Service.ScenesService.List()
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しないエラー")
	}
	scenesSize := len(scenes)
	if scenesSize == 0 {
		log.Println(err)
		response.BadRequest(writer, "不正なシナリオ名")
	}
	scenesList := make([]ResultScenesList, scenesSize, scenesSize)
	for i, v := range scenes {
		scenesList[i] = ResultScenesList{Name: v.Name, Num: v.Num, Action: v.Action, Text: v.Text, Image: v.Image}
	}
	response.Success(writer, ResponseScenesList{ScenesList: scenesList})
}

// ResponseScenesList /scene/listの返り値
type ResponseScenesList struct {
	ScenesList []ResultScenesList `json:"result"`
}

// ResultScenesList Sceneテーブルの返り値
type ResultScenesList struct {
	Name   string `json:"name"`
	Num    int    `json:"num"`
	Action string `json:"action"`
	Text   string `json:"text"`
	Image  string `json:"image"`
}
