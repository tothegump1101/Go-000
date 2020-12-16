package model

type User struct {
	Id uint64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age uint32 `json:"age" db:"age"`
}
