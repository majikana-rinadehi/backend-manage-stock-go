package usecases

import "github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"

type StockCategoryUsecase interface {
	GetAllStockCategories() ([]*entities.StockCategory, error)
}
