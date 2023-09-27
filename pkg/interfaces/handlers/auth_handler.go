package handlers

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	SignInWithProvider(c *gin.Context) *gin.Context
}
