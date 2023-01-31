package main

import (
	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/pkg/logger"
	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/pkg/repository"
)

func getRepo(log logger.Logger) repository.Repository {
	url, err := getDBUrl()
	if err != nil {
		log.Fatal(err.Error())
	}
	dbConn, err := repository.NewPostgreSqlConn(url)
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}
	return repository.NewPostgreSqlDB(dbConn)
}
