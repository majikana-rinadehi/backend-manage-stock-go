package usecases

import "github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"

type UserUsecase interface {
	GetAllUsers() ([]*entities.User, error)
}
