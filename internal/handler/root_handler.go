package handler

import (
	authHandler "clean_architecture/internal/service/login/deivery/http/handler"
)

type RootHandler struct {
	authHandler.AuthHandler
}

func NewRootHandler() *RootHandler {
	return &RootHandler{
		AuthHandler: authHandler.NewAuthHandler(),
	}
}
