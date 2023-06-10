package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// http.Handle("/foo", fooHandler)
	fmt.Println("run server")

	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "れいなだお, %q", html.EscapeString(r.URL.Path))
	// })

	connectMySQL()
	sqlInsert()
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func sqlInsert() {
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:rei0309@tcp(localhost:3306)/testrei")
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("INSERT INTO messages(message_text) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	// SQLの実行
	res, err := ins.Exec("こんばんは")
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
