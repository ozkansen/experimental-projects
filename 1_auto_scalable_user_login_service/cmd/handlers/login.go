package handlers

import (
	"net/http"

	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/pkg/encrypt"
	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/pkg/logger"
)

type Login struct {
	log           logger.Logger
	comparePasswd encrypt.Compare
}

func NewLoginHandler(c encrypt.Compare, l logger.Logger) *Login {
	return &Login{
		comparePasswd: c,
		log:           l,
	}
}

func (l *Login) Handler(w http.ResponseWriter, r *http.Request) {
}
