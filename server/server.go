package server

import (
	"log"
	"net/http"

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

	r := handler.NewRouter(config)
	http.ListenAndServe(port, r)

	return nil
}
