package dataSource

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig interface {
	getDSN() string
}

type PostgreSQL struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
}

//host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai

func (db PostgreSQL) getDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", db.Host, db.User, db.Password, db.DBName, db.Port)
}

func DBConnect(config DBConfig) (db *gorm.DB, err error) {
	dsn := config.getDSN()

	//后续有其他类型的数据库要支持就在这里加case
	switch config.(type) {
	case PostgreSQL:
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	return
}
