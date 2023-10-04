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

func (uc UserUsecase) GetUserAuths(userId int, authProvider string, uid string) ([]*entities.UserAuth, error) {
	userAuth, err := uc.r.FindAuth(userId, authProvider, uid)
	if err != nil {
		fmt.Println("Failed GetAllUser; ", err)
		return nil, err
	}
	return userAuth, nil
}

func (uc UserUsecase) SignInWithProvider(userId int, authProvider string, uid string) (*entities.User, error) {

	userAuths, err := uc.GetUserAuths(userId, authProvider, uid)
	if err != nil {
		fmt.Println("GetAllUsers failed", err)
		return nil, err
	}

	// すでにProviderに紐づくユーザが登録されていた場合
	if len(userAuths) > 0 {
		fmt.Println("すでにユーザが登録されています")
		users, err := uc.r.Find(userAuths[0].UserId)
		if err != nil {
			fmt.Println("GetAllUsers failed", err)
			return nil, err
		}
		return users[0], nil
	}

	newUser := &entities.User{
		Name:        "default name",
		MailAddress: "default email",
	}
	createdUser, err := uc.r.Save(newUser)
	if err != nil {
		fmt.Println("Failed CreateUser; ", err)
		return nil, err
	}

	newUserAuth := &entities.UserAuth{
		UserId:       createdUser.Id,
		AuthProvider: authProvider,
		Uid:          uid,
	}
	createdUserAuth, err := uc.r.SaveAuth(newUserAuth)
	if err != nil {
		fmt.Println("Failed CreateUserAuth; ", err)
		return nil, err
	}
	fmt.Println(createdUserAuth)

	return createdUser, nil
}

func (uc UserUsecase) UpdateUser(userId int, newUser *entities.User) error {

	err := uc.r.UpdateUser(userId, newUser)
	if err != nil {
		return err
	}

	return nil
}
