package databases

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDBConnection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
