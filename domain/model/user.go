package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            *int64 `bun:"id,pk" json:"id"`
	Name          string `bun:"name" json:"name"`
	Surname       string `bun:"surname" json:"surname"`
	Age           int64  `bun:"age" json:"age"`
}
