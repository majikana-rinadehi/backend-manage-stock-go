package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/majikana-rinadehi/backend-manage-stock-go/docs"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/handlers"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/usecases"
)

type StockHandler struct {
	stockUsecase usecases.StockUsecase
}

func NewStockHandler(uc usecases.StockUsecase) handlers.StockHandler {
	return &StockHandler{
		stockUsecase: uc,
	}
}

// GetAllStocks
// @Summary Stockを全件取得
// @Tags Stock
// @Produce json
// @Success 200 {array} entities.Stock
// @Failure 400
// @Failure 500
// @Router /stocks [get]
func (h *StockHandler) GetAllStocks(c *gin.Context) *gin.Context {
	stocks, err := h.stockUsecase.GetAllStocks()
	if err != nil {
		fmt.Println("GetAllStocks failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return c
	}
	c.JSON(http.StatusOK, stocks)
	return c
}

// CreateStock
// @Summary Stockを1件登録
// @Tags Stock
// @Produce json
// @Param body body entities.Stock false "Stock"
// @Success 200 {object} entities.Stock "登録したStock"
// @Failure 400
// @Failure 500
// @Router /stocks [post]
func (h *StockHandler) CreateStock(c *gin.Context) *gin.Context {
	var newStock entities.Stock

	if bindErr := c.BindJSON(&newStock); bindErr != nil {
		return c
	}

	stock, err := h.stockUsecase.CreateStock(&newStock)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return c
	}

	c.JSON(http.StatusOK, stock)
	return c
}

// DeleteStock
// @Summary idで指定したStockを1件削除する
// @Tags Stock
// @Produce json
// @Param id path string false "ID"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /stocks/{id} [delete]
func (h *StockHandler) DeleteStock(c *gin.Context) *gin.Context {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.stockUsecase.DeleteStock(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return c
	}

	c.Status(http.StatusNoContent)
	return c
}
