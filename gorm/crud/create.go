package crud

import (
	"golang-note/gorm/model"
	"gorm.io/gorm"
)

func Create(db *gorm.DB) {
	student := model.Student{
		Name:      "create test",
		Age:       0,
		IsDeleted: 0,
	}
	result := db.Create(&student)
	if result.Error != nil {
		panic("Failed to insert data, myError=" + result.Error.Error())
	}
}
