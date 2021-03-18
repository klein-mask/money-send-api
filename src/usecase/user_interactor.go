package usecase

import "money-send-api/domain"

type UserInteractor struct {
    UserRepository UserRepository
}

func (interactor *UserInteractor) AddUser(u domain.User) error {
    return interactor.UserRepository.AddUser(u)
}

func (interactor *UserInteractor) GetAllUsers() ([]domain.User, error) {
    return interactor.UserRepository.GetAllUsers()
}

func (interactor *UserInteractor) GetUser(userId string) (domain.User, error) {
    return interactor.UserRepository.GetUser(userId)
}

func (interactor *UserInteractor) UpdateAllBalance(balance int64) error {
    return interactor.UserRepository.UpdateAllBalance(balance)
}

func (interactor *UserInteractor) UpdateBalance(userId string, balance int64) error {
    return interactor.UserRepository.UpdateBalance(userId, balance)
}
/*
func (interactor *UserInteractor) DeleteAllUser() error {
    return interactor.UserRepository.DeleteAllUser()
}
*/
func (interactor *UserInteractor) DeleteUser(userId string) error {
    return interactor.UserRepository.DeleteUser(userId)
}
