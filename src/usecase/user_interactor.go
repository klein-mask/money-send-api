package usecase

import "money-send-api/src/domain"

type UserInteractor struct {
    UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) {
    interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) GetInfo() []domain.User {
    return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) SelectUser(id string) domain.User {
    return interactor.UserRepository.SelectUser(id)
}

func (interactor *UserInteractor) UpdateBalance(id string, balance int64) {
    interactor.UserRepository.UpdateBalance(id, balance)
}

func (interactor *UserInteractor) Delete(id string) {
    interactor.UserRepository.Delete(id)
}
