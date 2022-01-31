package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"os"
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

	log.Println(dataSource)
	conn, err := gorm.Open("mysql", dataSource)
	if err != nil {
		fmt.Println(err)
		return db, err
	}

	db.Conn = conn

	return db, nil
}

func InitMigration(db *gorm.DB) {
	for i := 1; i <= 7; i++ {
		fileName := fmt.Sprintf("config/database/migration%d.txt",i)
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		//exec.Command("mysql", "-u", "admin","-pP@ssword123","user_management","-e","source migration")
		if err := db.Exec(string(file)).Error; err!= nil {
			log.Fatal(err)
		}
	}
}