package usecases

import "github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"

type StockUsecase interface {
	GetAllStocks() ([]*entities.Stock, error)
	CreateStock(stock *entities.Stock) (*entities.Stock, error)
	DeleteStock(stockId int) error
	UpdateStock(stockId int, stock *entities.Stock) (*entities.Stock, error)
}
