package usecase

import "github.com/kmdavidds/election-api/internal/repository"

type Usecase struct {
	UserUsecase IUserUsecase
}

type InitParam struct {
	Repository *repository.Repository
}

func NewUsecase(param InitParam) *Usecase {
	userUsecase := NewUserUsecase(param)

	return &Usecase{
		UserUsecase: userUsecase,
	}
}
