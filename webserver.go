package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func main() {

	// ハンドラーを貼るパスの指定です。今回は他のディレクトリもないので、ルートを指定します。
	rootPath := "/"
	// 返却するファイルが配置されている場所のパスです。
	docDir := "./"

	http.HandleFunc(
		rootPath,
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, docDir)
		},
	)

	//log.Fatal(http.ListenAndServe(":8080", nil))
	appengine.Main()
}
