package entity

type Message struct {
	Id   string `db:"message_id"`
	Text string `db:"message_text"`
}
