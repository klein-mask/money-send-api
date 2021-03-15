package main

import (
    //"github.com/labstack/echo"
    //_ "github.com/lib/pq"
    //"gorm.io/gorm"
    //"money-send-api/src/domain"
    "money-send-api/src/infrastructure"
)
/*
var (
    db  *gorm.DB
    err error
    dsn = "root:password@tcp(127.0.0.1:3306)/go_sample?charset=utf8mb4&parseTime=True&loc=Local"
)
*/
func main() {
    //dbinit()
    infrastructure.Init()
    //e := echo.New()
    //e.Logger.Fatal(e.Start(":1323"))
}

/*
func dbinit() {
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
    }
    db.Migrator().CreateTable(domain.User{})
}
*/