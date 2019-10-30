package handler

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"prote-API/pkg/server/handler/response"
	"prote-API/pkg/server/service"
	"strconv"
	"time"
)

// GetHandleScenarioList /scenario/listのハンドラ(シーンリストの変更)
func GetHandleScenarioList(writer http.ResponseWriter, request *http.Request) {
	scenarios, err := service.Service.ScenarioService.List()
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しないエラー")
	}
	result := make([]ResultScenarioList, len(scenarios), len(scenarios))
	for i, v := range scenarios {
		result[i] = ResultScenarioList{Name: v.Name, Date: v.Date}
	}
	response.Success(writer, ResponseScenarioList{Result: result})
}

// PostHandleScenarioAdd
func PostHandleScenarioAdd(writer http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	dateS := request.FormValue("date")
	jst, _ := time.LoadLocation("Asia/Tokyo")
	date, err := time.ParseInLocation("2006-01-02 15:04:00", dateS, jst)
	if err != nil {
		response.BadRequest(writer, "不正な値")
	}
	err = service.Service.ScenarioService.Add(name, date)
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しないエラー")
	}
	response.Success(writer, "OK")
}

// PostHandleScenarioMake
func PostHandleScenarioMake(writer http.ResponseWriter, request *http.Request) {
	log.Println(request)
	name := request.FormValue("name")
	dateS := request.FormValue("date")
	jst, _ := time.LoadLocation("Asia/Tokyo")
	date, err := time.ParseInLocation("2006-01-02 15:04", dateS, jst)
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しない日付")
	}
	log.Println(name, date)

	sceneCnt, err := strconv.Atoi(request.FormValue("sceneCnt"))
	if err != nil {
		log.Println(err)
		response.BadRequest(writer, "予期しないシーン数")
	}
	var texts = make([]string, sceneCnt, sceneCnt)
	var acts = make([]string, sceneCnt, sceneCnt)
	var images = make([]string, sceneCnt, sceneCnt)
	for i := 0; i < sceneCnt; i++ {
		texts[i] = request.FormValue(fmt.Sprintf("text%v", i+1))
		acts[i] = request.FormValue(fmt.Sprintf("act%v", i+1))
		file, _, err := request.FormFile(fmt.Sprintf("image%v", i+1))
		if err != nil {
			if err == http.ErrMissingFile {
				log.Println(fmt.Sprintf("%v: ファイル無し", i+1))
			} else {
				log.Println(fmt.Sprintf("%v: ファイルエラー", i+1))
				log.Println(err)
				response.BadRequest(writer, "予期しないエラー")
			}
		} else {
			images[i] = fmt.Sprintf("%v_%v", name, i+1)
			buf := bytes.NewBuffer(nil)
			if _, err = io.Copy(buf, file); err != nil {
				log.Println(err)
			}
			var saveImage *os.File
			fileName := fmt.Sprintf("./image/%s.jpg", images[i])
			saveImage, e := os.Create(fileName)
			if e != nil {
				log.Println("サーバ側でファイル確保できませんでした。")
				return
			}
			_, err = io.Copy(saveImage, buf)
			file.Close()
		}
		log.Println(images[i])
	}
	err = service.Service.ScenarioService.Make(name, date, acts, texts, images)
	response.Success(writer, "OK")
}

/*
file, _, err := request.FormFile("image")
	name := request.FormValue("name")
	defer file.Close()
	if err != nil {
		log.Println(err)
	}
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, file); err != nil {
		log.Println(err)
	}

	var saveImage *os.File
	fileName := fmt.Sprintf("./image/%s.jpg", name)
	saveImage, e := os.Create(fileName)
	if e != nil {
		log.Println("サーバ側でファイル確保できませんでした。")
		return
	}
	_, err = io.Copy(saveImage, buf)
	response.Success(writer, "OK")
*/

type ResultScenarioList struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type ResponseScenarioList struct {
	Result []ResultScenarioList `json:"result"`
}
