package router

import (
	_ "github.com/Go-SQL-Driver/MySQL" 
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "net/http"
    "fmt"
    "time"
)

type User struct{
	UserId int `sql:"AUTO_INCREMENT"`
	Name string `sql:"size:32"`
	Password string `sql:"size:48"`
	Type int
	CreateTime time.Time
	ModifyTime time.Time 
}

var DB gorm.DB;

func init(){
	var err error
	if DB,err = gorm.Open("mysql", "root:1@/FishMoney?charset=utf8&parseTime=True") ; err != nil {
		fmt.Println("open mysql error!");
	}	
	fmt.Println(DB.DB());
	err = DB.DB().Ping()
	if err != nil {
		fmt.Println("open ping error!");
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func SetUserRoute(router *gin.RouterGroup){
	router.GET("/search", func(c *gin.Context) {
		var users []User;
		DB.Find(&users);
		fmt.Println(len(users));
		fmt.Println(users);
		c.String(http.StatusOK, "search");
	});
	router.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "get");
	});
	router.GET("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "add");
	});
}