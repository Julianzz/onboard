package model

import (
	"testing"

	"github.com/p1cn/onboard/liuzhenzhong/tests"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&DBTest{})

type DBTest struct {
	baseTest *tests.BaseTest
}

func (t *DBTest) SetUpTest(c *C) {
	t.baseTest = &tests.BaseTest{}
	t.baseTest.SetUpTest(c)
	DB = t.baseTest.DB
}

func (t *DBTest) TearDownTest(c *C) {
	t.baseTest.TearDownTest(c)
}
