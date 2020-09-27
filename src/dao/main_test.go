package dao_test

import (
	"os"
	"reserve_service/dao"
	"reserve_service/sys"
	"testing"
)

func TestMain(m *testing.M) {
	path, _ := os.Getwd()
	println(path)
	sys.InitConfig("../configs")
	dao.SetDriver()
	os.Exit(m.Run())
}
