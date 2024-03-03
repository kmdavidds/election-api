package usecase

import "github.com/kmdavidds/election-api/internal/repository"

type IUserUsecase interface {
}

type UserUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		ur: userRepository,
	}
}
