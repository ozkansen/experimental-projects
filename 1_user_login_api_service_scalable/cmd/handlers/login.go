package handlers

import (
	"context"
	"encoding/json"
	"io"
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

type Request struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}

func (l *Login) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		l.log.Warn("method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		l.log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var req Request
	if err = json.Unmarshal(data, &req); err != nil {
		l.log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbUser, err := l.repo.GetByUsername(context.Background(), req.Username)
	if err != nil {
		l.log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = l.comparePasswd.Compare([]byte(dbUser.Passwd), []byte(req.Passwd)); err != nil {
		l.log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte("ok"))
}
