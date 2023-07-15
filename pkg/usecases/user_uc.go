package usecases

import (
	"fmt"

	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/usecases"
)

type UserUsecase struct {
	r interfaces.UserRepository
}

func NewUserUsecase(r interfaces.UserRepository) usecases.UserUsecase {
	return &UserUsecase{
		r: r,
	}
}

func (uc UserUsecase) GetAllUsers() ([]*entities.User, error) {
	stocks, err := uc.r.FindAll()
	if err != nil {
		fmt.Println("Failed GetAllUser; ", err)
		return nil, err
	}
	return stocks, nil
}
