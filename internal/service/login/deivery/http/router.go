package http

import (
	handler "clean_architecture/internal/handler"
	//handler "clean_architecture/internal/service/login/deivery/http/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, rootHandler *handler.RootHandler) *gin.Engine {
	r.GET("/login", rootHandler.Login)
	return r
}
