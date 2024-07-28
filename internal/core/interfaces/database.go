package interfaces

import "gorm.io/gorm"

type Database interface {
	Create(value interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
}
