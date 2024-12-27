package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goodvs/internal/dao"
	"goodvs/internal/service"
	"goodvs/server"
	"net/http"
)

type Impl struct {
}

func (i Impl) Ping(c *gin.Context) {
	ResponseSuccess(c, "pong")
}

// Register register a new user
func (i Impl) Register(c *gin.Context) {
	//logrus.Info("Register")
	//fmt.Println("Register fmt")
	var user server.UserRegisterReq
	err := BindReq(c, &user)
	if err != nil {
		logrus.Info("BindReq failed", err)
		ResponseFail(c, err, http.StatusBadRequest)
		return
	}
	token, err := dao.DB(c).AddUser(user)
	if err != nil {
		logrus.Info("AddUser failed", err)
		ResponseFail(c, err, http.StatusBadRequest)
		return
	}
	ResponseSuccess(c, token)
}

// Login login
func (i Impl) Login(c *gin.Context) {
	var user server.UserLoginReq
	err := BindReq(c, &user)
	if err != nil {
		logrus.Info("BindReq failed", err)
		ResponseFail(c, err, http.StatusBadRequest)
		return
	}
	token, err := dao.DB(c).ValidateUser(user)
	if err != nil {
		logrus.Info("Login failed", err)
		ResponseFail(c, err, http.StatusBadRequest)
		return
	}
	ResponseSuccess(c, token)
}

// Search search
func (i Impl) Search(c *gin.Context) {
	// 参数仅为一个字符串
	reqStr, ok := c.GetQuery("product")
	if !ok {
		ResponseFail(c, fmt.Errorf("product is required"), http.StatusBadRequest)
		return
	}
	//fmt.Println(reqStr)
	//fmt.Println("Search")
	result, err := service.Search(reqStr)
	if err != nil {
		logrus.Error("Search failed: ", err) // 修改为错误日志级别
		ResponseFail(c, err, http.StatusBadRequest)
	}
	for _, v := range result.Results {
		fmt.Println(v)
	}
	ResponseSuccess(c, result)
}
