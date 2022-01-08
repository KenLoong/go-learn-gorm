package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint64
	Name      string
	Age       int8
	Birthday  sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 去除默认的表名复数
func (User) TableName() string {
	return "user"
}
