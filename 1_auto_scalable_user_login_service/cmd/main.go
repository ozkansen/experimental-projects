package main

import (
	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/pkg/logger"
)

func main() {
	log := logger.NewStdLogger(logger.InfoLevel)
	log.Info("starting user login api service")

	if err := Run(log); err != nil {
		log.Fatal(err.Error())
		return
	}
}

func Run(log logger.Logger) error {
	log.Info("running application service")
	return nil
}
