package infrastructure

import (
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
    "money-send-api/interfaces/database"
    "os"
)

type SqlHandler struct {
    db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
    host := os.Getenv("POSTGRES_HOST")
    if os.Getenv("IS_TEST") == "1" {
        host = os.Getenv("TEST_POSTGRES_HOST")
    }
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DB")
    port := os.Getenv("POSTGRES_PORT")

    dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
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

func (handler *SqlHandler) FindByName(obj interface{}, name string) error {
    return handler.db.Where("Name = ?", name).Find(obj).Error
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
