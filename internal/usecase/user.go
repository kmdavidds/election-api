package usecase

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/election-api/entity"
	"github.com/kmdavidds/election-api/internal/repository"
	"github.com/kmdavidds/election-api/model"
	"github.com/kmdavidds/election-api/pkg/bcrypt"
	"github.com/kmdavidds/election-api/pkg/jwt"
)

type IUserUsecase interface {
	Register(param model.UserRegister) error
	Login(param model.UserLogin) (model.UserLoginResponse, error)
	GetUser(param model.UserParam) (entity.User, error)
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
	hashedPassword, err := uu.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	param.ID = uuid.New()
	param.Password = hashedPassword

	user := entity.User{
		ID:       param.ID,
		Username: param.Username,
		Password: param.Password,
	}

	_, err = uu.ur.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) Login(param model.UserLogin) (model.UserLoginResponse, error) {
	result := model.UserLoginResponse{}

	user, err := uu.ur.GetUser(model.UserParam{
		Username: param.Username,
	})
	if err != nil {
		return result, err
	}

	err = uu.bcrypt.CompareHashAndPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := uu.jwt.CreateJWTToken(user.ID)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (uu *UserUsecase) GetUser(param model.UserParam) (entity.User, error) {
	return uu.ur.GetUser(param)
}