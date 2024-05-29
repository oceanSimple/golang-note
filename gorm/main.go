package main

import (
	"golang-note/gorm/crud"
	"golang-note/gorm/gorm"
)

var config = gorm.SqlConfig{
	Host:     "127.0.0.1",
	Port:     3306,
	Username: "root",
	Password: "root",
	DbName:   "project",
}

func main() {
	db := gorm.GetSqlConnect(config)
	crud.Create(db)
}
