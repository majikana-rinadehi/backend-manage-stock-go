package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/majikana-rinadehi/backend-manage-stock-go/docs"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/handlers"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/usecases"
)

type AuthHandler struct {
	userUsecase usecases.UserUsecase
}

func NewAuthHandler(uc usecases.UserUsecase) handlers.AuthHandler {
	return &AuthHandler{
		userUsecase: uc,
	}
}

// SignInWithProvider
// @Summary Providerによるサインインを実施する。新規ユーザの場合、ユーザ登録処理を実施する。
// @Tags Auth
// @Produce json
// @Param body body entities.UserAuth false "UserAuth"
// @Success 200 {array} entities.User
// @Failure 400
// @Failure 500
// @Router /auth/signin [post]
func (h *AuthHandler) SignInWithProvider(c *gin.Context) *gin.Context {

	var userAuth entities.UserAuth

	if bindErr := c.BindJSON(&userAuth); bindErr != nil {
		return c
	}

	// user_id
	// auth_provider
	// uid
	userId := userAuth.UserId
	authProvider := userAuth.AuthProvider
	uid := userAuth.Uid

	createdUser, err := h.userUsecase.SignInWithProvider(userId, authProvider, uid)
	if err != nil {
		fmt.Println("GetAllUsers failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return c
	}

	c.JSON(http.StatusOK, createdUser)
	return c
}
