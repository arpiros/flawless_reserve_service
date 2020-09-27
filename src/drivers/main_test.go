package drivers_test

import (
	"os"
	"reserve_service/sys"
	"testing"
)

func TestMain(m *testing.M) {
	path, _ := os.Getwd()
	println(path)
	sys.InitConfig("../configs")
	os.Exit(m.Run())
}
