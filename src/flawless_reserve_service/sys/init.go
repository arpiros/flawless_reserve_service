package sys

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func SystemInit() {
	InitConfig("./config")
}

func StartEchoFramework() {
	e := echo.New()
	e.File("/", "public/index.html")
	SetRouter(e)

	if viper.GetBool("use_tls") {
		e.Logger.Fatal(e.Start(viper.GetString("https_port")))
	} else {
		e.Logger.Fatal(e.Start(viper.GetString("http_port")))
	}
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
