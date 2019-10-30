package server

import (
	"log"
	"math/rand"
	"net/http"
	"prote-API/pkg/server/handler"
	"time"
)

// Serve HTTPサーバを起動する
func Serve(addr string) {
	rand.Seed(time.Now().UnixNano())
	/* ===== URLマッピングを行う ===== */
	http.HandleFunc("/test", get(handler.HandleTest))
	http.HandleFunc("/image/add", post(handler.PostHandleImageAdd))
	http.HandleFunc("/image/see", get(handler.GetHandleImageSee))
	http.HandleFunc("/image/get", get(handler.GetHandleImageGet))
	http.HandleFunc("/image/change", post(handler.PostHandleImageChange))
	http.HandleFunc("/scenes/list", get(handler.GetHandleScenesList))
	http.HandleFunc("/scene/list", get(handler.GetHandleSceneList))
	http.HandleFunc("/scene/add", post(handler.PostHandleSceneAdd))
	http.HandleFunc("/scene/delete", post(handler.PostHandleSceneDelete))
	http.HandleFunc("/scene/update", post(handler.PostHandleSceneUpdate))
	http.HandleFunc("/scenario/list", get(handler.GetHandleScenarioList))
	http.HandleFunc("/scenario/add", post(handler.PostHandleScenarioAdd))
	http.HandleFunc("/scenario/make", post(handler.PostHandleScenarioMake))
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		rand.Seed(time.Now().UnixNano())
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
