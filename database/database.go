package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tanirei39/test-rei/entity"
)

// MySQL接続確認
func Connect() {
	db := openDB()
	defer db.Close()
	// 実際に接続する
	err := db.Ping()
	if err != nil {
		log.Println("データベース接続失敗")
		log.Fatal(err)
	} else {
		log.Println("データベース接続完了！！")
	}
}

// Queryの実行
func SelectAll() []entity.Message {
	db := openDB()
	defer db.Close()
	ins, err := db.Prepare("SELECT message_id,message_text FROM messages")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	// クエリ実行
	rows, err := ins.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ml []entity.Message

	for rows.Next() {
		// 構造体宣言
		var message entity.Message
		err := rows.Scan(&message.Id, &message.Text)
		// 配列にScan結果を追加
		ml = append(ml, message)
		if err != nil {
			log.Println("データベース取得失敗")
			log.Fatal(err)
		}
	}

	return ml
}

func Select(id string) string {
	db := openDB()
	defer db.Close()

	// SQLの実行
	ins, err := db.Prepare("SELECT message_id,message_text FROM messages where message_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	// クエリ実行
	rows, err := ins.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var m entity.Message

	for rows.Next() {
		err := rows.Scan(&m.Id, &m.Text)

		if err != nil {
			log.Println("データベース取得失敗")
			log.Fatal(err)
		}
	}

	return m.Text
}

func Update(m entity.Message) {
	db := openDB()
	defer db.Close()

	// SQLの実行
	ins, err := db.Prepare("UPDATE messages set message_text = ?  where message_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	// クエリ実行
	res, err := ins.Exec(m.Text, m.Id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
}

func Insert(message string) {
	db := openDB()
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

func Delete(id string) {
	db := openDB()
	defer db.Close()

	// SQLの準備
	del, err := db.Prepare("DELETE FROM messages WHERE message_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer del.Close()

	// SQLの実行
	res, err := del.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	// 結果の取得
	_, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() *sql.DB {
	user := "root"
	password := "rei0309"
	host := "localhost"
	port := "3306"
	database_name := "testrei"

	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", dbconf)
	// db, err := sql.Open("mysql", "root:rei0309@tcp(localhost:3306)/testrei")

	if err != nil {
		log.Println("データベース接続失敗")
		log.Fatal(err)
	}
	return db
}
