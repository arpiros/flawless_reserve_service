package drivers_test

import (
	"reserve_service/sys"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	path, _ := os.Getwd()
	println(path)
	sys.InitConfig("../configs")
	m.Run()
}
