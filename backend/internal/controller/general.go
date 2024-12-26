package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"net/http"
)

// >>>>>>>>>>>>>>> Response >>>>>>>>>>>>>>>>>>

type Resp struct {
	Code int `json:"code" example:"200"` // status code
	Data any `json:"data"`               // data payload
}

func Response(ctx *gin.Context, code int, data any) {
	if _, ok := data.(string); ok {
		ctx.String(code, data.(string))
	} else {
		ctx.JSON(code, data)
	}
}

func ResponseSuccess(ctx *gin.Context, data any) {
	Response(ctx, http.StatusOK, data)
}

func ResponseFail(ctx *gin.Context, err error) {
	logrus.Error(err)
}

// <<<<<<<<<<<<<<<<< Response >>>>>>>>>>>>>>>>>

// >>>>>>>>>>>>>>>> Request >>>>>>>>>>>>>>>>>>

// BindReq bind request, and log info&err.
// Support json & query & header,
// if there is a conflict, the former will overwrite the latter.
func BindReq[T any](c *gin.Context, req *T) error {
	logrus.Debugf("Req.Url: %s, Req.Body: %+v", c.Request.URL, c.Request.Body)
	if err := c.ShouldBindWith(req, GeneralBinder); err != nil {
		logrus.Errorf("BindReq failed: %s", err)
	}
	return nil
}

var (
	GeneralBinder = General{}
)

type General struct{}

func (General) Name() string {
	return "general"
}

func (General) Bind(r *http.Request, obj any) error {
	// general binder
	// 1. bind header
	// 2. bind query or form
	// 3. bind body (json)
	_ = binding.Header.Bind(r, obj)
	_ = binding.Query.Bind(r, obj)
	_ = binding.JSON.Bind(r, obj)
	return binding.Validator.ValidateStruct(obj)
}

// <<<<<<<<<<<<<<<<< Request <<<<<<<<<<<<<<<<<<
