package controllers

import (
	"net/http"
)

// この二つの引数を持つものはハンドラーとして登録される
func top(w http.ResponseWriter, r *http.Request) {

	// // ParseFilesでhtmlファイルを解析している
	// t, err := template.ParseFiles("app/views/templates/top.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // 解析されたtを実行、第2引数はhtmlに渡すパラメータ
	// // t.Execute(w, nil)
	// t.Execute(w, "Hello")

	generateHTML(w, "Hello", "layout", "top")
}
