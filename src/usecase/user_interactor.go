package usecase

import "money-send-api/src/domain"

type UserInteractor struct {
    UserRepository UserRepository
}

func (interactor *UserInteractor) AddUser(u domain.User) {
    interactor.UserRepository.AddUser(u)
}

func (interactor *UserInteractor) GetAllUsers() []domain.User {
    return interactor.UserRepository.GetAllUsers()
}

func (interactor *UserInteractor) GetUser(id string) domain.User {
    return interactor.UserRepository.GetUser(id)
}

func (interactor *UserInteractor) UpdateAllBalance(balance int64) {
    interactor.UserRepository.UpdateAllBalance(balance)
}

func (interactor *UserInteractor) UpdateBalance(id string, balance int64) {
    interactor.UserRepository.UpdateBalance(id, balance)
}

func (interactor *UserInteractor) DeleteUser(id string) {
    interactor.UserRepository.DeleteUser(id)
}
