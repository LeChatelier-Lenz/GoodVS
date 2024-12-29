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
	_, err = dao.DB(c).AddUser(user)
	if err != nil {
		logrus.Info("AddUser failed", err)
		ResponseFail(c, err, http.StatusBadRequest)
		return
	}
	ResponseSuccess(c, nil)
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

// Search general search
func (i Impl) Search(c *gin.Context) {
	// 参数仅为一个字符串
	reqStr, ok := c.GetQuery("product")
	if !ok {
		ResponseFail(c, fmt.Errorf("product is required"), http.StatusBadRequest)
		return
	}
	result, err := service.SearchCallByFrontend(reqStr)
	if err != nil {
		logrus.Error("Search failed: ", err) // 修改为错误日志级别
		ResponseFail(c, err, http.StatusBadRequest)
	}
	if len(result) == 0 {
		ResponseFail(c, fmt.Errorf("no result searched"), http.StatusBadRequest)
		return
	}
	errCount := 0
	for _, v := range result {
		productId, err := dao.DB(c).AddProductItem(v)
		if err != nil {
			logrus.Error("AddProductItem failed: ", err) // 修改为错误日志级别
			errCount++
			continue
		}
		err = dao.DB(c).AddProductPrice(productId, v.Price)
		if err != nil {
			logrus.Error("AddProductPrice failed: ", err) // 修改为错误日志级别
			errCount++
			continue
		}
	}
	if errCount >= len(result)/2 {
		// 保存失败的数据超过一半
		Response(c, http.StatusOK, server.SearchRes{
			Msg:     "search success, but some data failed to save",
			Results: result,
		})
		return
	}
	for _, v := range result {
		fmt.Println(v)
	}
	ResponseSuccess(c, server.SearchRes{
		Msg:     "search success",
		Results: result,
	})
}

// PlatformLogin platform login
func (i Impl) PlatformLogin(c *gin.Context) {
	var req server.PlatformLoginReq
	err := BindReq(c, &req)
	if err != nil {
		ResponseFail(c, err, http.StatusBadRequest)
		return
	}
	platform := req.Platform
	if platform == "" {
		ResponseFail(c, fmt.Errorf("platform is required"), http.StatusBadRequest)
		return
	}
	err = service.PlatformLogin(c, platform)
	if err != nil {
		ResponseFail(c, err, http.StatusBadRequest)
		return
	}
	ResponseSuccess(c, "login "+platform+" success")
}

// Follow follow a product
func (i Impl) Follow(c *gin.Context) {
	var req server.FollowReq
	err := BindReq(c, &req)
	if err != nil {
		ResponseFail(c, err, http.StatusBadRequest)
		return
	}
	err = dao.DB(c).AddFollow(req.ProductId, req.UserId)
	if err != nil {
		ResponseFail(c, err, http.StatusBadRequest)
		return
	}
	ResponseSuccess(c, "follow success")
}

// Unfollow unfollow a product
func (i Impl) Unfollow(c *gin.Context) {
	var req server.FollowReq
	err := BindReq(c, &req)
	if err != nil {
		ResponseFail(c, err, http.StatusBadRequest)
		return
	}
	err = dao.DB(c).RemoveFollow(req.ProductId, req.UserId)
	if err != nil {
		ResponseFail(c, err, http.StatusBadRequest)
		return
	}
	ResponseSuccess(c, "unfollow success")
}
