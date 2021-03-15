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

func (interactor *UserInteractor) GetUser() domain.User {
    return interactor.UserRepository.SelectUser()
}

func (interactor *UserInteractor) Delete(id string) {
    interactor.UserRepository.Delete(id)
}
