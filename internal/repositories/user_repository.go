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

func (r *UserRepository) FindAll(userId int, authProvider string, uid string) (users []*entities.UserAuth, err error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	chain := db.Where("")

	if userId != 0 {
		chain.Where("user_id = ?", userId)
	}

	if authProvider != "" {
		chain.Where("auth_provider = ?", authProvider)
	}

	if uid != "" {
		chain.Where("uid = ?", uid)
	}

	if err := chain.Debug().Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Save(user *entities.User) (*entities.User, error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) SaveAuth(userAuth *entities.UserAuth) (*entities.UserAuth, error) {
	db, dbErr := r.dbAdapter.GetDB()
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.Create(&userAuth).Error; err != nil {
		return nil, err
	}

	return userAuth, nil
}
