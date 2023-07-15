package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/majikana-rinadehi/backend-manage-stock-go/docs"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/handlers"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/usecases"
)

type StockCategoryHandler struct {
	stockCategoryUsecase usecases.StockCategoryUsecase
}

func NewStockCategoryHandler(uc usecases.StockCategoryUsecase) handlers.StockCategoryHandler {
	return &StockCategoryHandler{
		stockCategoryUsecase: uc,
	}
}

// GetAllCategories
// @Summary Categoryを全件取得
// @Tags Category
// @Produce json
// @Success 200 {array} entities.StockCategory
// @Failure 400
// @Failure 500
// @Router /categories [get]
func (h *StockCategoryHandler) GetAllCategories(c *gin.Context) *gin.Context {
	stocks, err := h.stockCategoryUsecase.GetAllStockCategories()
	if err != nil {
		fmt.Println("GetAllStocks failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return c
	}
	c.JSON(http.StatusOK, stocks)
	return c
}
