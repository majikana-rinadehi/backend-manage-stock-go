package usecases

import "github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"

type UserUsecase interface {
	GetUserAuths(userId int, authProvider string, uid string) ([]*entities.UserAuth, error)
	SignInWithProvider(userId int, authProvider string, uid string) (*entities.User, error)
	UpdateUser(userId int, newUser *entities.User) error
}
