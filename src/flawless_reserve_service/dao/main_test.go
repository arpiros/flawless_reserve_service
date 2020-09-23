package dao_test

import (
	"flawless_reserve_service/flawless_reserve_service/dao"
	"flawless_reserve_service/flawless_reserve_service/sys"
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
