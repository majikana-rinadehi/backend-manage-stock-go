package handlers

import (
	"github.com/gin-gonic/gin"
)

type StockCategoryHandler interface {
	GetCategories(c *gin.Context) *gin.Context
	CreateCategory(c *gin.Context) *gin.Context
	DeleteCategory(c *gin.Context) *gin.Context
	UpdateCategory(c *gin.Context) *gin.Context
}
