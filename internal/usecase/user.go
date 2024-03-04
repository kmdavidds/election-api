package usecase

import (
	"github.com/kmdavidds/election-api/internal/repository"
	"github.com/kmdavidds/election-api/model"
	"github.com/kmdavidds/election-api/pkg/bcrypt"
	"github.com/kmdavidds/election-api/pkg/jwt"
)

type IUserUsecase interface {
	Register(param model.UserRegister) error
}

type UserUsecase struct {
	ur     repository.IUserRepository
	bcrypt bcrypt.Interface
	jwt    jwt.Interface
}

func NewUserUsecase(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwt jwt.Interface) IUserUsecase {
	return &UserUsecase{
		ur:     userRepository,
		bcrypt: bcrypt,
		jwt:    jwt,
	}
}

func (uu *UserUsecase) Register(param model.UserRegister) error {
	return nil
}
