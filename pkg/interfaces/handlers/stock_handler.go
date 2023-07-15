package handlers

import "github.com/gin-gonic/gin"

type StockHandler interface {
	GetAllStocks(c *gin.Context) *gin.Context
	CreateStock(c *gin.Context) *gin.Context
	DeleteStock(c *gin.Context) *gin.Context
}
