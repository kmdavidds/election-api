package middleware

import (
	"github.com/kmdavidds/election-api/internal/usecase"
)

type Interface interface {
}

type middleware struct {
	usecase *usecase.Usecase
}

func Init(usecase *usecase.Usecase) Interface {
	return &middleware{
		usecase: usecase,
	}
}
