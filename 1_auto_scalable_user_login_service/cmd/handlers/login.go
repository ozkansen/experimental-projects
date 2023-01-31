package handlers

import (
	"net/http"

	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/pkg/logger"
)

type Login struct {
	log logger.Logger
}

func NewLoginHandler(l logger.Logger) *Login {
	return &Login{log: l}
}

func (l *Login) Handler(w http.ResponseWriter, r *http.Request) {
}
