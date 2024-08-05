package db

import (
	"gorm.io/gorm"
)

var GormDB *gorm.DB

type Student struct {
	gorm.Model
	Code  string `gorm:"not null;unique" binding:"required"`
	Image string `gorm:"not null"`
	Elo   int    `gorm:"default:1400"`
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(&Student{})
}
