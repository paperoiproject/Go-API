package handler

import (
	"fmt"
	"log"
	"net/http"
	"prote-API/pkg/server/handler/response"
	"prote-API/pkg/server/service"
	"strconv"
)

// GetHandleSceneList /scene/listのハンドラ(シーンリストの変更)
func GetHandleSceneList(writer http.ResponseWriter, request *http.Request) {
	Form := TestRequest{
		Name: request.FormValue("name"),
	}
	scenes, err := service.Service.SceneService.List(Form.Name)
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しないエラー")
	}
	scenesSize := len(scenes)
	if scenesSize == 0 {
		log.Println(err)
		response.BadRequest(writer, "不正なシナリオ名")
	}
	sceneList := make([]ResultSceneList, scenesSize, scenesSize)
	for i, v := range scenes {
		sceneList[i] = ResultSceneList{Name: v.Name, Num: v.Num, Action: v.Action, Text: v.Text}
	}
	response.Success(writer, ResponseSceneList{SceneList: sceneList})
}

// ResponseSceneList /scene/listの返り値
type ResponseSceneList struct {
	SceneList []ResultSceneList `json:"scene_list"`
}

// ResultSceneList Sceneテーブルの返り値
type ResultSceneList struct {
	Name   string `json:"name"`
	Num    int    `json:"num"`
	Action string `json:"action"`
	Text   string `json:"text"`
}

// PostHandleSceneAdd /scene/addのハンドラ(シーンリストの変更)
func PostHandleSceneAdd(writer http.ResponseWriter, request *http.Request) {
	FormNum, err := strconv.Atoi(request.FormValue("num"))
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "不正な値")
	}
	FormName := request.FormValue("name")
	var Tasks = make([]Task, FormNum, FormNum)
	var works = make([]string, FormNum, FormNum)
	var texts = make([]string, FormNum, FormNum)

	for i := 0; i < FormNum; i++ {
		work := request.FormValue(fmt.Sprintf("work%v", i+1))
		text := request.FormValue(fmt.Sprintf("text%v", i+1))
		Tasks[i] = Task{
			Work: work,
			Text: text,
		}
		works[i] = work
		texts[i] = text
	}
	err = service.Service.SceneService.Add(FormName, FormNum, works, texts)
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しないエラー")
	}
	response.Success(writer, ResponseSceneAdd{Name: FormName, Tasks: Tasks})
}

type Task struct {
	Work string `json:"work"`
	Text string `json:"text"`
}

type ResponseSceneAdd struct {
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

// PostHandleSceneDelete /scene/deleteのハンドラ(シーンリストの変更)
func PostHandleSceneDelete(writer http.ResponseWriter, request *http.Request) {
	FormName := request.FormValue("name")
	err := service.Service.SceneService.Delete(FormName)
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しないエラー")
	}
	response.Success(writer, "OK")
}

// PostHandleSceneUpdate /scene/updateのハンドラ(シーンリストの変更)
func PostHandleSceneUpdate(writer http.ResponseWriter, request *http.Request) {
	FormNum, err := strconv.Atoi(request.FormValue("num"))
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "不正な値")
	}
	FormName := request.FormValue("name")
	var Tasks = make([]Task, FormNum, FormNum)
	var works = make([]string, FormNum, FormNum)
	var texts = make([]string, FormNum, FormNum)
	for i := 0; i < FormNum; i++ {
		work := request.FormValue(fmt.Sprintf("work%v", i+1))
		text := request.FormValue(fmt.Sprintf("text%v", i+1))
		Tasks[i] = Task{
			Work: work,
			Text: text,
		}
		works[i] = work
		texts[i] = text
	}

	err = service.Service.SceneService.Update(FormName, FormNum, works, texts)

	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しないエラー")
	}
	response.Success(writer, ResponseSceneAdd{Name: FormName, Tasks: Tasks})
}
