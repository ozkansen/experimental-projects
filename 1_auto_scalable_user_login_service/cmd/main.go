package main

import (
	"net/http"
	"time"

	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/cmd/handlers"
	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/pkg/encrypt"
	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/pkg/logger"
)

func main() {
	log := logger.NewStdLogger(logger.InfoLevel)
	log.Info("starting user login api service")

	passwdCompare := encrypt.NewBcryptCompare()

	loginHandler := handlers.NewLoginHandler(passwdCompare, log)
	mux := router(loginHandler)

	log.Info("running application service")
	if err := Run(mux); err != nil {
		log.Fatal(err.Error())
		return
	}
}

func Run(mux *http.ServeMux) error {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
