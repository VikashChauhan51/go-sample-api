package interfaces

import "gorm.io/gorm"

type Database interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}
