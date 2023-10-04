package usecases

import "github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"

type StockCategoryUsecase interface {
	GetStockCategories(userId int) ([]*entities.StockCategory, error)
	CreateCategory(category *entities.StockCategory) (*entities.StockCategory, error)
	DeleteCategory(categoryId int) error
	UpdateCategory(stockId int, stock *entities.StockCategory) (*entities.StockCategory, error)
}
