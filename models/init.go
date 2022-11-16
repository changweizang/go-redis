package models

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitMysql() {
	dsn := "root:123456@tcp(106.15.199.75:3308)/hmdp?charset=utf8mb4&parseTime=True&loc=Local"
	dbClient, _ :=sql.Open("mysql", dsn)
	dbCli, err := gorm.Open(mysql.New(mysql.Config{
		Conn: dbClient,
	}))
	if err != nil {
		log.Fatalln(err)
	}
	db = dbCli
	db = db.Debug()
}

func GetDb() *gorm.DB {
	return db
}


