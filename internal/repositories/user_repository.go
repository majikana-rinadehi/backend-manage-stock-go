package repositories

import (
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/adapters"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces"
)

type UserRepository struct {
	dbAdapter *adapters.DatabaseAdapter
}

func NewUserRepository(a *adapters.DatabaseAdapter) interfaces.UserRepository {
	return &UserRepository{
		dbAdapter: a,
	}
}

func (r *UserRepository) FindAll() (stocks []*entities.User, err error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.Find(&stocks).Error; err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *UserRepository) Save(stock *entities.User) (*entities.User, error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.Create(&stock).Error; err != nil {
		return nil, err
	}

	return stock, nil
}
