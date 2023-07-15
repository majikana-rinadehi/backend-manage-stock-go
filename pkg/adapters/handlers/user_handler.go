package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/majikana-rinadehi/backend-manage-stock-go/docs"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/handlers"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/usecases"
)

type UserHandler struct {
	stockUsecase usecases.UserUsecase
}

func NewUserHandler(uc usecases.UserUsecase) handlers.UserHandler {
	return &UserHandler{
		stockUsecase: uc,
	}
}

// GetAllUsers
// @Summary Userを全件取得
// @Tags User
// @Produce json
// @Success 200 {array} entities.User
// @Failure 400
// @Failure 500
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) *gin.Context {
	users, err := h.stockUsecase.GetAllUsers()
	if err != nil {
		fmt.Println("GetAllUsers failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return c
	}
	c.JSON(http.StatusOK, users)
	return c
}
