package handler

import (
	"github.com/gorilla/mux"
	"github.com/p1cn/onboard/liuzhenzhong/config"
)

// NewRouter hander route path
func NewRouter(config *config.Config) *mux.Router {

	r := mux.NewRouter()

	userHandler := &RestfulHandler{
		H: &UsersHandler{},
	}
	r.Handle("/users", userHandler).Methods("GET", "POST")

	relationHandler := &RestfulHandler{
		H: &RelationsHandler{},
	}
	r.Handle("/users/{user_id}/relationships", relationHandler).Methods("GET")
	r.Handle("/users/{user_id}/relationships/{wipe_user_id}", relationHandler).Methods("PUT")

	return r
}
