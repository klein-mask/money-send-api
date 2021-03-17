package database

type SqlHandler interface {
    Create(object interface{})
    FindById(object interface{}, id string)
    FindAll(object interface{})
    Update(obj interface{}, cond string, condValue interface{}, column string, columnValue interface{})
    UpdateByExpr(obj interface{}, cond string, condValue interface{}, column string, columnExpr string, columnValue ...interface{})
    DeleteById(object interface{}, id string)
}
