package handler

import (
	usecase "clean_architecture/internal/service/login/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authnHandler struct {
	usecase usecase.AuthUseCase
	//service service.Service
}

type AuthHandler interface {
	Login(c *gin.Context)
	//Logout(c *gin.Context)
	//GetSignMessage(c *gin.Context)
	//LoginWithMetamask(c *gin.Context)
	//LogUserLocation(c *gin.Context)
}

func NewAuthHandler() AuthHandler {
	return &authnHandler{
		usecase: usecase.NewAuthUseCase(),
	}
}

func (a *authnHandler) Login(c *gin.Context) {
	a.usecase.Login()
	c.JSONP(http.StatusOK, gin.H{
		"data": 123,
	})
}
