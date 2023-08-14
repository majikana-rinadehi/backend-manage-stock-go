package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/go-ozzo/ozzo-validation/v4"
	_ "github.com/majikana-rinadehi/backend-manage-stock-go/docs"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/handlers"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/usecases"
	"github.com/majikana-rinadehi/backend-manage-stock-go/util"
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
		c.JSON(http.StatusInternalServerError, &handlers.Response[*entities.Stock]{
			Total:   0,
			Results: nil,
			Errors: []*handlers.ErrorResponse{
				{
					Message: "GetAllStocks failed",
				},
			},
		})
		return c
	}
	c.JSON(http.StatusOK, &handlers.Response[*entities.Stock]{
		Total:   len(stocks),
		Results: stocks,
		Errors:  nil,
	})
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
		c.JSON(http.StatusBadRequest, &handlers.Response[any]{
			Total:   0,
			Results: nil,
			Errors: []*handlers.ErrorResponse{
				{
					Message: bindErr.Error(),
				},
			},
		})
		return c
	}

	// リクエストバリデーションチェック
	vErr := validateStock(newStock)

	if vErr != nil {
		handlers.BadRequests(c, vErr.(validation.Errors))
		return c
	}

	stock, err := h.stockUsecase.CreateStock(&newStock)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &handlers.Response[*entities.Stock]{
			Total:   0,
			Results: nil,
			Errors: []*handlers.ErrorResponse{
				{
					Message: "CreateStock failed",
				},
			},
		})
		return c
	}

	c.JSON(http.StatusOK, &handlers.Response[*entities.Stock]{
		Total: 1,
		Results: []*entities.Stock{
			stock,
		},
		Errors: nil,
	})
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
		c.JSON(http.StatusInternalServerError, &handlers.Response[any]{
			Total:   0,
			Results: nil,
			Errors: []*handlers.ErrorResponse{
				{
					Message: "DeleteStock failed",
				},
			},
		})
		return c
	}

	c.Status(http.StatusNoContent)
	return c
}

// UpdateStock
// @Summary idで指定したStockを1件更新する
// @Tags Stock
// @Produce json
// @Param id path string false "ID"
// @Param body body entities.Stock false "Stock"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /stocks/{id} [put]
func (h *StockHandler) UpdateStock(c *gin.Context) *gin.Context {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)

	var newStock entities.Stock

	if bindErr := c.BindJSON(&newStock); bindErr != nil {
		c.JSON(http.StatusBadRequest, &handlers.Response[any]{
			Total:   0,
			Results: nil,
			Errors: []*handlers.ErrorResponse{
				{
					Message: bindErr.Error(),
				},
			},
		})
		return c
	}

	updatedStock, err := h.stockUsecase.UpdateStock(id, &newStock)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &handlers.Response[any]{
			Total:   0,
			Results: nil,
			Errors: []*handlers.ErrorResponse{
				{
					Message: "UpdateStock failed",
				},
			},
		})
		return c
	}

	c.JSON(http.StatusOK, &handlers.Response[*entities.Stock]{
		Total: 1,
		Results: []*entities.Stock{
			updatedStock,
		},
		Errors: nil,
	})
	return c
}

// validateStock validates Stock data in request body
func validateStock(newStock entities.Stock) error {

	// リクエストバリデーションチェック
	vErr := validation.ValidateStruct(&newStock,
		validation.Field(&newStock.UserId,
			validation.By(util.ValidateIntNotEmpty("userId")),
		),
		validation.Field(&newStock.CategoryId,
			validation.By(util.ValidateIntNotEmpty("categoryId")),
		),
		validation.Field(&newStock.Name,
			validation.By(util.ValidateStrNotEmpty("name")),
			validation.Length(0, 255).Error(util.MaxLengthErrMsg("name", 255)),
		),
		validation.Field(&newStock.Amount,
			validation.By(util.ValidateIntNotEmpty("amount")),
		),
		validation.Field(&newStock.ExpireDate,
			validation.By(util.ValidateYYYY_MM_DD("expireDate")),
		),
	)

	return vErr
}
