package handlers

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetUserAuths(c *gin.Context) *gin.Context
	PutUser(c *gin.Context) *gin.Context
}
