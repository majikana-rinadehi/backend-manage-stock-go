package repositories

import (
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/adapters"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces"
	"gorm.io/gorm/clause"
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

func (r *StockCategoryRepository) FindByQuery(userId int) (stockCategories []*entities.StockCategory, err error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	chain := db.Where("")

	if userId != 0 {
		chain.Where("user_id = ?", userId)
	}

	if err := chain.Debug().Find(&stockCategories).Error; err != nil {
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

func (r *StockCategoryRepository) DeleteById(id int) error {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return dbErr
	}

	if err := db.
		Debug().
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Delete(&entities.StockCategory{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *StockCategoryRepository) Update(id int, category *entities.StockCategory) (*entities.StockCategory, error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.
		Model(&entities.StockCategory{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"name": category.Name,
		}).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *StockCategoryRepository) CountById(id int) (int64, error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return 0, dbErr
	}

	var count int64

	if err := db.
		Model(&entities.StockCategory{}).
		Where("id = ?", id).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
