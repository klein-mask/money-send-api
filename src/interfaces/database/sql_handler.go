package database

type SqlHandler interface {
    Create(object interface{}) error
    FindAll(object interface{}) error
    FindById(object interface{}, id string) error
    Update(obj interface{}, cond string, condValue interface{}, column string, columnValue interface{}) error
    UpdateByExpr(obj interface{}, cond string, condValue interface{}, column string, columnExpr string, columnValue ...interface{}) error
    DeleteById(object interface{}, id string) error
}
