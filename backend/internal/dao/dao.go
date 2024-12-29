package dao

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"goodvs/internal/dao/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
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
	// "user:pass@tcp(127.0.0.1:3306)/dbname"
	User   string `mapstructure:"User"`
	Pass   string `mapstructure:"Pass"`
	Host   string `mapstructure:"Host"`
	Port   string `mapstructure:"Port"`
	DBName string `mapstructure:"DBName"`
}

func (cfg DBCfg) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DBName)
}

func InitDB() {
	var cfg DBCfg
	err := viper.Sub("Database").UnmarshalExact(&cfg)
	if err != nil {
		fmt.Println("Error unmarshalling database config")
		fmt.Println(err)
	}

	db, err = gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{
		TranslateError:                           true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("Error opening database connection")
		fmt.Println(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Product{}, &model.ProductPrice{}, &model.Follow{})
	if err != nil {
		fmt.Println("Error migrating database")
		fmt.Println(err)
	}

	// 初始化定时查询服务
	Interval := 1 * time.Minute

	timerService := NewTimerService(Interval, DBMS{db})

	// 启动定时查询服务
	timerService.Start()

	if viper.GetString("App.RunLevel") == "debug" {
		db = db.Debug()
	}
	fmt.Println("Database connection established")

}
