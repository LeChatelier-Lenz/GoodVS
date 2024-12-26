package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"price-compare/server"
	"time"
)

type WebServerCfg struct {
	Port         int `mapstructure:"Port"`
	WriteTimeout int `mapstructure:"WriteTimeout"`
	ReadTimeout  int `mapstructure:"ReadTimeout"`
}

func StartServer() error {
	var cfg WebServerCfg
	if err := viper.Sub("WebServer").UnmarshalExact(&cfg); err != nil {
		return err
	}

	e := gin.Default()

	// 添加CORS中间件
	e.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://chiral-phonon-material-database.vercel.app/")
		// 注意上面的header中的域名需要和前端的域名一致，否则前端会报跨域错误，需要修改
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	server.RegisterHandlers(e.Group("/api"), Impl{})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        e,
		ReadTimeout:    time.Second * time.Duration(cfg.ReadTimeout),
		WriteTimeout:   time.Second * time.Duration(cfg.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}
