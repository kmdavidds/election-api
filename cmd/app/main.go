package main

import (
	"github.com/kmdavidds/election-api/internal/delivery/rest"
	"github.com/kmdavidds/election-api/internal/repository"
	"github.com/kmdavidds/election-api/internal/usecase"
	"github.com/kmdavidds/election-api/pkg/config"
	"github.com/kmdavidds/election-api/pkg/database/mysql"
	"github.com/kmdavidds/election-api/pkg/middleware"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	db := mysql.ConnectDatabase()

	mysql.Migrate(db)

	repository := repository.NewRepository(db)

	usecase := usecase.NewUsecase(usecase.InitParam{Repository: repository})

	middleware := middleware.Init(usecase)

	rest := rest.NewRest(usecase, middleware)

	rest.MountEndpoint()

	rest.Serve()
}
