package repositories

import (
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/adapters"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces"
	"gorm.io/gorm/clause"
)

type StockRepository struct {
	dbAdapter *adapters.DatabaseAdapter
}

func NewStockRepository(a *adapters.DatabaseAdapter) interfaces.StockRepository {
	return &StockRepository{
		dbAdapter: a,
	}
}

func (r *StockRepository) FindAll() (stocks []*entities.Stock, err error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.Find(&stocks).Error; err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *StockRepository) FindByQuery(userId int) (stocks []*entities.Stock, err error) {

	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	chain := db.Where("")

	if userId != 0 {
		chain.Where("user_id = ?", userId)
	}

	if err := chain.Debug().Find(&stocks).Error; err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *StockRepository) Save(stock *entities.Stock) (*entities.Stock, error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.Create(&stock).Error; err != nil {
		return nil, err
	}

	return stock, nil
}

func (r *StockRepository) DeleteById(id int) error {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return dbErr
	}

	if err := db.
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Delete(&entities.Stock{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *StockRepository) Update(id int, stock *entities.Stock) (*entities.Stock, error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.
		Model(&entities.Stock{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"category_id": stock.CategoryId,
			"name":        stock.Name,
			"amount":      stock.Amount,
			"expire_date": stock.ExpireDate,
		}).Error; err != nil {
		return nil, err
	}

	return stock, nil
}

func (r *StockRepository) CountById(id int) (int64, error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return 0, dbErr
	}

	var count int64

	if err := db.
		Model(&entities.Stock{}).
		Where("id = ?", id).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
