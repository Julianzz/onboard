package handler

import (
	"log"
	"net/http"
	"testing"

	"github.com/p1cn/onboard/liuzhenzhong/tests"

	"github.com/p1cn/onboard/liuzhenzhong/context"

	"github.com/go-pg/pg"
	. "gopkg.in/check.v1"

	"github.com/p1cn/onboard/liuzhenzhong/config"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&HandlerTest{})

type HandlerTest struct {
	db     *pg.DB
	server *http.Server
	port   string

	baseTest *tests.BaseTest
}

func (t *HandlerTest) SetUpTest(c *C) {
	t.baseTest = &tests.BaseTest{}

	t.baseTest.SetUpTest(c)

	conf := &config.Config{
		DBSetting: config.DBSetting{
			Host:     "127.0.0.1:5432",
			User:     "liuzhenzhong",
			Database: "test",
			Password: "",
		},
	}
	t.port = "127.0.0.1:9090"
	context.NewServerContext(conf)

	r := NewRouter(conf)

	srv := &http.Server{
		Addr:    t.port,
		Handler: r,
	}
	t.server = srv

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

}

func (t *HandlerTest) TearDownTest(c *C) {
	t.baseTest.TearDownTest(c)
	if t.server != nil {
		t.server.Close()
	}
}
