package handlers

import "github.com/gin-gonic/gin"

type StockCategoryHandler interface {
	GetAllCategories(c *gin.Context) *gin.Context
}
