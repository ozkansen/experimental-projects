package handlers

import (
	"net/http"

	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/pkg/encrypt"
	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/pkg/logger"
	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/pkg/repository"
)

type Login struct {
	log           logger.Logger
	comparePasswd encrypt.Compare
	repo          repository.Repository
}

func NewLoginHandler(r repository.Repository, c encrypt.Compare, l logger.Logger) *Login {
	return &Login{
		repo:          r,
		comparePasswd: c,
		log:           l,
	}
}

func (l *Login) Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok\n"))
}
