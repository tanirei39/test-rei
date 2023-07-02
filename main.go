package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tanirei39/test-rei/database"
)

func main() {
	fmt.Println("run server")

	database.Connect()
	http.HandleFunc("/message-form", HandlerMessageForm)
	http.HandleFunc("/message-confirm", RedirectHandlerMessageConfirm)
	http.HandleFunc("/message-list", HandlerUserConfirm)
	http.HandleFunc("/message-delete/", HandlerMessageDelete)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

// 投稿入力画面
func HandlerMessageForm(w http.ResponseWriter, r *http.Request) {
	// テンプレートに出力する値をマップにセット
	values := map[string]string{}

	// マップを展開してテンプレートを出力&テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/message-form.gtpl"))
	if err := tpl.ExecuteTemplate(w, "message-form.gtpl", values); err != nil {
		fmt.Println(err)
	}
}

// 　投稿登録画面(リダイレクトで一覧に戻る)
func RedirectHandlerMessageConfirm(w http.ResponseWriter, r *http.Request) {
	//データ登録
	database.Insert(r.FormValue("message_text"))
	http.Redirect(w, r, "/message-list", 301)
}

// 投稿一覧画面
func HandlerUserConfirm(w http.ResponseWriter, r *http.Request) {
	// MySQLからデータ取得(テンプレートに出力する値)
	ml := database.Select("SELECT message_id,message_text FROM messages")

	// マップを展開してテンプレートを出力&テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/message-list.gtpl"))
	if err := tpl.ExecuteTemplate(w, "message-list.gtpl", ml); err != nil {
		fmt.Println(err)
	}
}

// 　投稿削除画面(リダイレクトで一覧に戻る)
func HandlerMessageDelete(w http.ResponseWriter, r *http.Request) {
	//id取得&データ削除
	id := strings.TrimPrefix(r.URL.Path, "/message-delete/")
	database.Delete(id)
	//一覧画面へリダイレクト
	http.Redirect(w, r, "/message-list", 301)
}
