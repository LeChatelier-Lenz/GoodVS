// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// (GET /ping)
	Ping(c *gin.Context)
	// (GET /user/register)
	Register(c *gin.Context)
	// (GET /user/login)
	Login(c *gin.Context)
	// (GET /search
	Search(c *gin.Context)
	// (GET /platform/login)
	PlatformLogin(c *gin.Context)
	// (POST /follow)
	Follow(c *gin.Context)
	// (POST /unfollow)
	Unfollow(c *gin.Context)
	// (GET /follow)
	GetFollowList(c *gin.Context)
	// (GET /pricelist)
	GetPriceList(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)




// Ping operation middleware
func (siw *ServerInterfaceWrapper) Ping(c *gin.Context) {
	fmt.Println("ping")
	c.String(http.StatusOK, "pong")
}

// Register operation middleware
func (siw *ServerInterfaceWrapper) Register(c *gin.Context) {
	fmt.Println("Process: user register")
	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}
	siw.Handler.Register(c)
}

// Login operation middleware
func (siw *ServerInterfaceWrapper) Login(c *gin.Context) {
	fmt.Println("Process: user login")
	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}
	siw.Handler.Login(c)
}

// Search operation middleware
func (siw *ServerInterfaceWrapper) Search(c *gin.Context) {
	fmt.Println("Process: search")
	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}
	siw.Handler.Search(c)
}

// PlatformLogin operation middleware
func (siw *ServerInterfaceWrapper) PlatformLogin(c *gin.Context) {
	fmt.Println("Process: platform login")
	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}
	siw.Handler.PlatformLogin(c)
}

// Follow operation middleware
func (siw *ServerInterfaceWrapper) Follow(c *gin.Context) {
	fmt.Println("Process: follow")
	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}
	siw.Handler.Follow(c)
}

// Unfollow operation middleware
func (siw *ServerInterfaceWrapper) Unfollow(c *gin.Context) {
	fmt.Println("Process: unfollow")
	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}
	siw.Handler.Unfollow(c)
}


// GetFollowList operation middleware
func (siw *ServerInterfaceWrapper) GetFollowList(c *gin.Context) {
	fmt.Println("Process: get follow list")
	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}
	siw.Handler.GetFollowList(c)
}

// GetPriceList operation middleware
func (siw *ServerInterfaceWrapper) GetPriceList(c *gin.Context) {
	fmt.Println("Process: get price list")
	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}
	siw.Handler.GetPriceList(c)
}


// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/ping", wrapper.Ping)
	router.POST(options.BaseURL+"/user/register", wrapper.Register)
	router.POST(options.BaseURL+"/user/login", wrapper.Login)
	router.GET(options.BaseURL+"/search", wrapper.Search)
	router.POST(options.BaseURL+"/platform/login",wrapper.PlatformLogin)
	router.POST(options.BaseURL+ "/follow", wrapper.Follow)
	router.POST(options.BaseURL+ "/unfollow", wrapper.Unfollow)
	router.GET(options.BaseURL+"/follow", wrapper.GetFollowList)
	router.GET(options.BaseURL+"/pricelist", wrapper.GetPriceList)
	// 专门用于处理非法路由请求
}
