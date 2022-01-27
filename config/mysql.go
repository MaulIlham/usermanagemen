package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Conn *gorm.DB
}

func InitDB(keyUser, keyPassword, keyHost, keyPort, keySchema string) (Database, error) {
	db := Database{}
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		os.Getenv(keyUser),
		os.Getenv(keyPassword),
		os.Getenv(keyHost),
		os.Getenv(keyPort),
		os.Getenv(keySchema),
	)

	conn, err := gorm.Open("mysql", dataSource)
	if err != nil {
		fmt.Println(err)
		return db, err
	}

	db.Conn = conn

	return db, nil
}
