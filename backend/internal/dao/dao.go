package dao

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBMS struct {
	*gorm.DB
}

var (
	db *gorm.DB
)

var DB = func(ctx context.Context) *DBMS {
	return &DBMS{db.WithContext(ctx)}

}

// >>>>>>>>>>>> init >>>>>>>>>>>>

type DBCfg struct {
	DSN string
}

func InitDB() {
	var cfg DBCfg
	err := viper.Sub("Database").UnmarshalExact(&cfg)
	if err != nil {
		fmt.Println("Error unmarshalling database config")
		fmt.Println(err)
	}

	db, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		fmt.Println("Error opening database connection")
		fmt.Println(err)
	}

	if viper.GetString("App.RunLevel") == "debug" {
		db = db.Debug()
	}

}
