// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"sample/pkg/api"
	"sample/pkg/api/handler"
	"sample/pkg/config"
	"sample/pkg/db"
	"sample/pkg/repository"
	"sample/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg *config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	
	userRepository := repository.NewUserDataBase(gormDB)
	userUseCase := usecase.NewuserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	adminRepository:= repository.NewAdminRepository(gormDB)
	adminUseCase:= usecase.NewAdminUseCase(adminRepository)
	adminHandler:=handler.NewAdminHandler(adminUseCase)


	serverHTTP := http.NewServerHttp(userHandler,adminHandler)
	return serverHTTP, nil
}