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

type StockCategoryHandler struct {
	stockCategoryUsecase usecases.StockCategoryUsecase
}

func NewStockCategoryHandler(uc usecases.StockCategoryUsecase) handlers.StockCategoryHandler {
	return &StockCategoryHandler{
		stockCategoryUsecase: uc,
	}
}

// GetCategories
// @Summary 条件に一致するCategoryを取得
// @Tags Category
// @Produce json
// @Param userId query string false "userId"
// @Success 200 {array} entities.StockCategory
// @Failure 400
// @Failure 500
// @Router /categories [get]
func (h *StockCategoryHandler) GetCategories(c *gin.Context) *gin.Context {

	userId, _ := strconv.Atoi(c.Query("userId"))

	categories, err := h.stockCategoryUsecase.GetStockCategories(userId)
	if err != nil {
		fmt.Println("GetAllStocks failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return c
	}
	c.JSON(http.StatusOK, &handlers.Response[*entities.StockCategory]{
		Total:   len(categories),
		Results: categories,
		Errors:  nil,
	})
	return c
}

// CreateCategory
// @Summary Categoryを1件登録
// @Tags Category
// @Produce json
// @Param body body entities.StockCategory false "Category"
// @Success 200 {object} entities.StockCategory "登録したCategory"
// @Failure 400
// @Failure 500
// @Router /categories [post]
func (h *StockCategoryHandler) CreateCategory(c *gin.Context) *gin.Context {
	var newCategory entities.StockCategory

	if bindErr := c.BindJSON(&newCategory); bindErr != nil {
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
	// vErr := validateStock(newCategory)

	// if vErr != nil {
	// 	handlers.BadRequests(c, vErr.(validation.Errors))
	// 	return c
	// }

	category, err := h.stockCategoryUsecase.CreateCategory(&newCategory)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &handlers.Response[*entities.StockCategory]{
			Total:   0,
			Results: nil,
			Errors: []*handlers.ErrorResponse{
				{
					Message: "CreateCategory failed",
				},
			},
		})
		return c
	}

	c.JSON(http.StatusOK, &handlers.Response[*entities.StockCategory]{
		Total: 1,
		Results: []*entities.StockCategory{
			category,
		},
		Errors: nil,
	})
	return c
}

// DeleteCategory
// @Summary idで指定したCategoryを1件削除する
// @Tags Category
// @Produce json
// @Param id path string false "ID"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /categories/{id} [delete]
func (h *StockCategoryHandler) DeleteCategory(c *gin.Context) *gin.Context {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.stockCategoryUsecase.DeleteCategory(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &handlers.Response[any]{
			Total:   0,
			Results: nil,
			Errors: []*handlers.ErrorResponse{
				{
					Message: "DeleteCategory failed",
				},
			},
		})
		return c
	}

	c.Status(http.StatusNoContent)
	return c
}

// UpdateCategory
// @Summary idで指定したCategoryを1件更新する
// @Tags Category
// @Produce json
// @Param id path string false "ID"
// @Param body body entities.StockCategory false "Category"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /categories/{id} [put]
func (h *StockCategoryHandler) UpdateCategory(c *gin.Context) *gin.Context {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)

	var newCategory entities.StockCategory

	if bindErr := c.BindJSON(&newCategory); bindErr != nil {
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

	updatedCategory, err := h.stockCategoryUsecase.UpdateCategory(id, &newCategory)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &handlers.Response[any]{
			Total:   0,
			Results: nil,
			Errors: []*handlers.ErrorResponse{
				{
					Message: "UpdateCategory failed",
				},
			},
		})
		return c
	}

	c.JSON(http.StatusOK, &handlers.Response[*entities.StockCategory]{
		Total: 1,
		Results: []*entities.StockCategory{
			updatedCategory,
		},
		Errors: nil,
	})
	return c
}
