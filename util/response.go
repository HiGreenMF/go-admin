package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

const (
	Success       = 0
	Error         = -1
	InvalidParams = 401
	Unauthorized  = 403
)

func (g *Gin) Response(httpCode, code int, data any,msg string) {
	g.C.JSON(httpCode, ResponseData{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	return
}

func (g *Gin) ResponseOk(code int,data any, msg string) {
	g.Response(http.StatusOK, code, data, msg)
}

func (g *Gin) ResponseCreated(code int,data any, msg string) {
	g.Response(http.StatusCreated, code, data, msg)
}

func (g *Gin) ResponseAccepted(code int,data any, msg string) {
	g.Response(http.StatusAccepted, code, data, msg)
}

func (g *Gin) ResponseBadRequest(code int,data any, msg string) {
	g.Response(http.StatusBadRequest, code, data, msg)
}

func (g *Gin) ResponseUnauthorized(code int,data any, msg string) {
	g.Response(http.StatusUnauthorized, code, data, msg)
}

func (g *Gin) ResponseForbidden(code int,data any, msg string) {
	g.Response(http.StatusForbidden, code, data, msg)
}

func (g *Gin) ResponseNotFound(code int,data any, msg string) {
	g.Response(http.StatusNotFound, code, data, msg)
}

func (g *Gin) ResponseServerError(code int,data any, msg string) {
	g.Response(http.StatusInternalServerError, code, data, msg)
}
