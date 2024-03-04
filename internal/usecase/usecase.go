package usecase

import (
	"github.com/kmdavidds/election-api/internal/repository"
	"github.com/kmdavidds/election-api/pkg/bcrypt"
	"github.com/kmdavidds/election-api/pkg/jwt"
)

type Usecase struct {
	UserUsecase IUserUsecase
}

type InitParam struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.Interface
	JWT        jwt.Interface
}

func NewUsecase(param InitParam) *Usecase {
	userUsecase := NewUserUsecase(param.Repository.UserRepository, param.Bcrypt, param.JWT)

	return &Usecase{
		UserUsecase: userUsecase,
	}
}
