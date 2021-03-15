package infrastructure

import (
	"database/sql"
	_ "github.com/lib/pq" 
    "gorm.io/gorm"

    "echoSample/src/interfaces/database"
)

type SqlHandler struct {
    db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
    dsn := "root:password@tcp(127.0.0.1:3306)/go_sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("postgres", "host=postgres user=admin password=adminpass dbname=app sslmode=disable")
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err.Error)
    }
    sqlHandler := new(SqlHandler)
    sqlHandler.db = db
    return sqlHandler
}

func (handler *SqlHandler) Create(obj interface{}) {
    handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
    handler.db.Find(obj)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
    handler.db.Delete(obj, id)
}
