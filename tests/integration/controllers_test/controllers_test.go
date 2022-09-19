package controllers_test

import (
	"clean_gin_api/tests/setup"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	code := m.Run()
	setup.TeardownDB()
	os.Exit(code)
}
