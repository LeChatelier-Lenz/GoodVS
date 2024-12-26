package controller

import "github.com/gin-gonic/gin"

type Impl struct {
}

func (i Impl) Ping(c *gin.Context) {
	ResponseSuccess(c, "pong")
}
