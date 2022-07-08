package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	host := "localhost"
	port := "3306"
	dbname := "myapp_db"
	username := "root"
	password := ""

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	DB, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction: true,
		})

	if err != nil {
		return nil, err
	}
	return DB, nil
}
