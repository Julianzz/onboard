package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/p1cn/onboard/liuzhenzhong/config"
	"github.com/p1cn/onboard/liuzhenzhong/context"
	"github.com/p1cn/onboard/liuzhenzhong/handler"
)

// StartServer init server context, start server
func StartServer(port string, config *config.Config) error {

	serverContext, err := context.NewServerContext(config)
	if err != nil {
		log.Panicln("error in init server context")
		return err
	}
	context.Context = serverContext

	r := mux.NewRouter()

	userHandler := &handler.RestfulHandler{
		H: &handler.UsersHandler{},
	}
	r.Handle("/users", userHandler).Methods("GET", "POST")

	relationHandler := &handler.RestfulHandler{
		H: &handler.RelationsHandler{},
	}
	r.Handle("/users/{user_id}/relationships", relationHandler).Methods("GET")
	r.Handle("/users/{user_id}/relationships/{wipe_user_id}", relationHandler).Methods("PUT")

	http.ListenAndServe(port, r)

	return nil
}
