package handler

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"os"
	"prote-API/pkg/server/handler/response"
)

var viewImg = "a"

// PostHandleImageAdd image/addのハンドラ(画像の追加)
func PostHandleImageAdd(writer http.ResponseWriter, request *http.Request) {
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
}

// GetHandleImageSee  /image/seeのハンドラ(imageの返却)
func GetHandleImageSee(writer http.ResponseWriter, request *http.Request) {
	file, err := os.Open(fmt.Sprintf("./image/%s.jpg", viewImg))
	defer file.Close()
	if err != nil {
		log.Println(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
	}
	log.Println("OK")
	response.SuccessFile(writer, &img)
}

// GetHandleImageGet  /image/getのハンドラ(imageの返却)
func GetHandleImageGet(writer http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	file, err := os.Open(fmt.Sprintf("./image/%s.jpg", name))
	defer file.Close()
	if err != nil {
		log.Println(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
	}
	log.Println("OK")
	response.SuccessFile(writer, &img)
}

// PostHandleImageChange  /image/changeのハンドラ(imageの返却)
func PostHandleImageChange(writer http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	viewImg = name
	response.Success(writer, "OK")
}
