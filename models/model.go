package models

type User struct {
	ID    uint `json:"id"`
	Name  string
	Email string
}

type Session struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	BLOB   []byte `json:"blob"`
}
