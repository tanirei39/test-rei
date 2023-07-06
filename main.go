package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tanirei39/test-rei/database"
	"github.com/tanirei39/test-rei/entity"
)

type MessageForm struct {
	Title1   string
	Title2   string
	Message  string
	Action   string
	DoneText string
}

var url string = "http://localhost:8080/"

func main() {
	fmt.Println("run server")

	database.Connect()
	http.HandleFunc("/message-form", HandlerMessageForm)
	http.HandleFunc("/message-confirm", RedirectHandlerMessageConfirm)
	http.HandleFunc("/message-list", HandlerUserConfirm)
	http.HandleFunc("/message-delete/", HandlerMessageDelete)
	http.HandleFunc("/message-edit/", HandlerMessageEdit)
	http.HandleFunc("/message-edited/", HandlerMessageEdited)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

// 投稿一覧画面
func HandlerUserConfirm(w http.ResponseWriter, r *http.Request) {
	// MySQLからデータ取得(テンプレートに出力する値)
	ml := database.SelectAll()

	// マップを展開してテンプレートを出力&テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/message-list.gtpl"))
	if err := tpl.ExecuteTemplate(w, "message-list.gtpl", ml); err != nil {
		fmt.Println(err)
	}
}

// 投稿入力画面
func HandlerMessageForm(w http.ResponseWriter, r *http.Request) {
	// テンプレートに出力する値をマップにセット
	values := MessageForm{Title1: "新規投稿", Title2: "新規投稿", Message: "", Action: url + "message-confirm", DoneText: "投稿"}

	// マップを展開してテンプレートを出力&テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/message-form.gtpl"))
	if err := tpl.ExecuteTemplate(w, "message-form.gtpl", values); err != nil {
		fmt.Println(err)
	}
}

// 　投稿編集画面
func HandlerMessageEdit(w http.ResponseWriter, r *http.Request) {
	//id取得&データ削除
	id := strings.TrimPrefix(r.URL.Path, "/message-edit/")
	log.Println(id)
	m := database.Select(id)
	log.Println(m)
	// テンプレートに出力する値をマップにセット
	values := MessageForm{Title1: "投稿編集", Title2: "No" + id + "投稿編集", Message: m, Action: url + "message-edited/" + id, DoneText: "更新"}

	// マップを展開してテンプレートを出力&テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/message-form.gtpl"))
	if err := tpl.ExecuteTemplate(w, "message-form.gtpl", values); err != nil {
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

// 　投稿登録画面(リダイレクトで一覧に戻る)
func RedirectHandlerMessageConfirm(w http.ResponseWriter, r *http.Request) {
	//データ登録
	log.Println(r.FormValue("message_text"))
	database.Insert(r.FormValue("message_text"))
	http.Redirect(w, r, "/message-list", 301)
}

// 　投稿編集完了画面(リダイレクトで一覧に戻る)
func HandlerMessageEdited(w http.ResponseWriter, r *http.Request) {
	//データ更新
	log.Println(r.FormValue("message_text"))
	id := strings.TrimPrefix(r.URL.Path, "/message-edited/")
	m := entity.Message{Id: id, Text: r.FormValue("message_text")}
	database.Update(m)

	http.Redirect(w, r, "/message-list", 301)
}
