package context

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/p1cn/onboard/liuzhenzhong/config"
	"github.com/p1cn/onboard/liuzhenzhong/model"
)

// ServerContext server context contain the db connection
type ServerContext struct {
	Config *config.Config
	Db     *pg.DB
}

// Context for init common module
var Context *ServerContext

// NewServerContext server context info
func NewServerContext(config *config.Config) (*ServerContext, error) {
	// init db connection
	db, err := model.InitDB(config)
	if err != nil {
		log.Printf("error in initing db error:%v", err)
		return nil, err
	}

	context := &ServerContext{
		Config: config,
		Db:     db,
	}
	return context, nil
}
