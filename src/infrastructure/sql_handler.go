package infrastructure

import (
	"gorm.io/driver/postgres"
    "gorm.io/gorm"

    "money-send-api/interfaces/database"
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

func (handler *SqlHandler) Create(obj interface{}) error {
    return handler.db.Create(obj).Error
}

func (handler *SqlHandler) FindAll(obj interface{}) error {
    return handler.db.Find(obj).Error
}

func (handler *SqlHandler) FindById(obj interface{}, id string) error {
    return handler.db.Where("ID = ?", id).Find(obj).Error
}

func (handler *SqlHandler) Update(obj interface{}, cond string, condValue interface{}, column string, columnValue interface{}) error {
    return handler.db.Model(obj).Where(cond, condValue).Update(column, columnValue).Error
}

func (handler *SqlHandler) UpdateByExpr(obj interface{}, cond string, condValue interface{}, column string, columnExpr string, columnValue ...interface{}) error {
    return handler.db.Model(obj).Where(cond, condValue).Update(column, gorm.Expr(columnExpr, columnValue)).Error
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) error {    
    tx := handler.db.Begin()
    defer func() {
      if r := recover(); r != nil {
        tx.Rollback()
      }
    }()
  
    if err := tx.Error; err != nil {
      return err
    }
  
    if err := tx.Where("id = ?", id).Delete(obj).Error; err != nil {
       tx.Rollback()
       return err
    }
  
    if err := tx.Unscoped().Where("id = ?", id).Delete(obj).Error; err != nil {
       tx.Rollback()
       return err
    }
  
    return tx.Commit().Error
}
