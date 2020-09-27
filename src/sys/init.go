package sys

import (
	"reserve_service/dao"
	"reserve_service/middleware"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func SystemInit() {
	InitConfig("./configs")
	dao.SetDriver()
}

func StartEchoFramework() {
	e := echo.New()
	e.File("/", "public/index.html")
	e.Debug = viper.GetBool("debug_mode")

	middleware.SetMiddleWare(e)
	SetRouter(e)

	e.Logger.Fatal(e.Start(viper.GetString("http_port")))
}

func InitConfig(path string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	if viper.InConfig("include") {
		includeName := viper.GetString("include")
		viper.SetConfigName(includeName)
		if err := viper.ReadInConfig(); err != nil {
			panic(err.Error())
		}
	}
}
