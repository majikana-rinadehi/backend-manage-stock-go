package repositories

import (
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/adapters"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces"
)

type StockCategoryRepository struct {
	dbAdapter *adapters.DatabaseAdapter
}

func NewStockCategoryRepository(a *adapters.DatabaseAdapter) interfaces.StockCategoryRepository {
	return &StockCategoryRepository{
		dbAdapter: a,
	}
}

func (r *StockCategoryRepository) FindAll() (stockCategories []*entities.StockCategory, err error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.Find(&stockCategories).Error; err != nil {
		return nil, err
	}

	return stockCategories, nil
}

func (r *StockCategoryRepository) Save(stockCategory *entities.StockCategory) (*entities.StockCategory, error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.Create(&stockCategory).Error; err != nil {
		return nil, err
	}

	return stockCategory, nil
}
