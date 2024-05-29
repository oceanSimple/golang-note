package model

import "time"

type Student struct {
	Id          uint64 `gorm:"primaryKey"`
	Name        string
	Age         uint64
	IsDeleted   uint8
	GmtCreate   time.Time
	GmtModified time.Time
}
