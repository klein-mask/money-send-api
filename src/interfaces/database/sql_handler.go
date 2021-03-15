package database

type SqlHandler interface {
    Create(object interface{})
    FindAll(object interface{})
    FindUser(object interface{})
    DeleteById(object interface{}, id string)
}
