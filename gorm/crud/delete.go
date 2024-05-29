package crud

import (
	"golang-note/gorm/model"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB) {
	var stu = model.Student{Id: 1}
	db.Debug().Delete(&stu)
}
