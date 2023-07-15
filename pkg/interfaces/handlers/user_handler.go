package handlers

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetAllUsers(c *gin.Context) *gin.Context
}
