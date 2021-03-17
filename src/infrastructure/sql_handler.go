package infrastructure

import (
	"gorm.io/driver/postgres"
    "gorm.io/gorm"

    "money-send-api/src/interfaces/database"
)

type SqlHandler struct {
    db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
    dsn := "host=postgres user=admin password=admin_pass dbname=app port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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

func (handler *SqlHandler) FindById(obj interface{}, id string) {
    handler.db.Where("ID = ?", id).Find(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
    handler.db.Find(obj)
}

func (handler *SqlHandler) Update(obj interface{}, cond string, condValue interface{}, column string, columnValue interface{}) {
    handler.db.Model(obj).Where(cond, condValue).Update(column, columnValue)
}

func (handler *SqlHandler) UpdateByExpr(obj interface{}, cond string, condValue interface{}, column string, columnExpr string, columnValue ...interface{}) {
    handler.db.Model(obj).Where(cond, condValue).Update(column, gorm.Expr(columnExpr, columnValue))
}


func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
    handler.db.Delete(obj, id)
}
