package interfaces

import (
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
)

type StockRepository interface {
	Save(stock *entities.Stock) (*entities.Stock, error)
	FindAll() ([]*entities.Stock, error)
	DeleteById(id int) error
	Update(id int, stock *entities.Stock) (*entities.Stock, error)
	CountById(id int) (int64, error)
}

type StockCategoryRepository interface {
	Save(stock *entities.StockCategory) (*entities.StockCategory, error)
	FindAll() ([]*entities.StockCategory, error)
}

type UserRepository interface {
	Save(user *entities.User) (*entities.User, error)
	SaveAuth(userAuth *entities.UserAuth) (*entities.UserAuth, error)
	FindAll(userId int, authProvider string, uid string) ([]*entities.UserAuth, error)
}
