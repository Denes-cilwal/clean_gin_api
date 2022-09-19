package service_test

import (
	"clean_gin_api/tests/setup"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	setup.TeardownDB()
	os.Exit(code)
}
