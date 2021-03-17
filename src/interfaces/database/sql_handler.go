package database

// database/user_repositoryが引数として受けるために必要
type SqlHandler interface {
    Create(object interface{})
    //FindUser(object interface{})
    FindById(object interface{}, id string)
    FindAll(object interface{})
    UpdateBalanceById(obj interface{}, id string, balance int64)
    DeleteById(object interface{}, id string)
}
