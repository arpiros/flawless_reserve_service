package dao_test

import (
	"reserve_service/dao"
	"reserve_service/sys"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	path, _ := os.Getwd()
	println(path)
	sys.InitConfig("../configs")
	dao.SetDriver()
	m.Run()
}
