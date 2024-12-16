package models

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"name"`
	PassHash []byte `json:"pass_hash"`
}
