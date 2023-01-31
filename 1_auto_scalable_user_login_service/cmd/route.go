package main

import (
	"net/http"

	"github.com/ozkansen/experimental-projects/1_auto_scalable_user_login_service/cmd/handlers"
)

func router(login *handlers.Login) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", login.Handler)
	return mux
}
