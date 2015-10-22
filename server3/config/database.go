package config

import (
	_ "github.com/Go-SQL-Driver/MySQL" 
    "github.com/jinzhu/gorm"
    "fmt"
)

var DB *gorm.DB;

func init(){
	tempDb,err := gorm.Open("mysql", "root:1@/FishMoney?charset=utf8&parseTime=True") 
	if err != nil {
		fmt.Println("open mysql error!");
	}
	DB = &tempDb;
	//DB.LogMode(true)
	err = DB.DB().Ping()
	if err != nil {
		fmt.Println("open ping error!");
	}
	DB.DB().SetMaxIdleConns(10000)
	DB.DB().SetMaxOpenConns(10000)
}