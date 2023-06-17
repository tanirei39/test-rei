package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// マスタからSELECTしたデータをマッピングする構造体
type MESSAGE struct {
	Id   string `db:"message_id"`
	Text string `db:"message_text"`
}

func main() {
	// http.Handle("/foo", fooHandler)
	fmt.Println("run server")

	// http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "れいなだお, %q", html.EscapeString(r.URL.Path))
	// })

	connectMySQL()
	// "user-form"へのリクエストを関数で処理する
	http.HandleFunc("/message-form", HandlerUserForm)
	// "message-confirm"へのリクエストを関数で処理する
	http.HandleFunc("/message-confirm", RedirectHandlerMessageConfirm)
	// "user-confirm"へのリクエストを関数で処理する
	http.HandleFunc("/message-list", HandlerUserConfirm)

	// sqlInsert()
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// 投稿入力画面
func HandlerUserForm(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/message-form.gtpl"))

	// テンプレートに出力する値をマップにセット
	values := map[string]string{}

	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "message-form.gtpl", values); err != nil {
		fmt.Println(err)
	}
}

func RedirectHandlerMessageConfirm(w http.ResponseWriter, r *http.Request) {
	//データ登録
	messageInsert(r.FormValue("message_text"))
	http.Redirect(w, r, "/message-list", 301)
}

// 投稿一覧画面
func HandlerUserConfirm(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/message-list.gtpl"))

	// データベース接続
	db, err := sql.Open("mysql", "root:rei0309@tcp(localhost:3306)/testrei")
	if err != nil {
		// log.Fatal(err)
		fmt.Println("データベース接続に失敗しました。")
	}
	// deferで処理終了前に必ず接続をクローズする
	defer db.Close()

	// プリペアードステートメント
	ins, err := db.Prepare("SELECT message_id,message_text FROM messages")
	if err != nil {
		log.Fatal(err)
	}

	// クエリ実行
	rows, err := ins.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ml []MESSAGE

	for rows.Next() {
		// 構造体宣言
		var message MESSAGE
		err = rows.Scan(&message.Id, &message.Text)
		// 配列にScan結果を追加
		ml = append(ml, message)
	}

	if err != nil {
		fmt.Println("データベース接続失敗")
		log.Fatal(err)
	} else {
		fmt.Println("データベース接続成功")
	}

	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "message-list.gtpl", ml); err != nil {
		fmt.Println(err)
	}
}

func messageInsert(message string) {
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:rei0309@tcp(localhost:3306)/testrei")
	if err != nil {
		// ここではエラーを返さない
		fmt.Println("データベース接続に失敗しました。")
	}
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("INSERT INTO messages(message_text) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	// SQLの実行
	res, err := ins.Exec(message)
	if err != nil {
		log.Fatal(err)
	}

	// 結果の取得
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(lastInsertID)
}

func connectMySQL() {
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:rei0309@tcp(localhost:3306)/testrei")
	// db, err := sql.Open("mysql", "gouser:hogehoge@tcp(localhost:3306)/go_test")
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	defer db.Close()

	// 実際に接続する
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("データベース接続完了")
	}
}
