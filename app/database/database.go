package database

import "github.com/jinzhu/gorm"

type Database interface {
	Kill()
	Migrate(tables []interface{})
	Init() *gorm.DB
}
