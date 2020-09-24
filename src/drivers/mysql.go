package drivers

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() *gorm.DB {
	dbInfo := viper.GetString("common_db")
	db, err := gorm.Open(
		mysql.Open(dbInfo),
		&gorm.Config{
			Logger: NewLogger(),
		},
	)
	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	sqlDB.SetMaxIdleConns(viper.GetInt("max_idle_conns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("max_open_conns"))

	return db
}

func NewLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
}
