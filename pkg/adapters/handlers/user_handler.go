package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

// GetUserAuths
// @Summary 検索条件に合うUserを取得
// @Tags User
// @Produce json
// @Param userId query string false "userId"
// @Param authProvider query string false "authProvider"
// @Param uid query string false "uid"
// @Success 200 {array} entities.UserAuth
// @Failure 400
// @Failure 500
// @Router /user-auths [get]
func (h *UserHandler) GetUserAuths(c *gin.Context) *gin.Context {

	// user_id
	// auth_provider
	// uid
	userId, _ := strconv.Atoi(c.Query("userId"))
	authProvider := c.Query("authProvider")
	uid := c.Query("uid")

	users, err := h.stockUsecase.GetUserAuths(userId, authProvider, uid)
	if err != nil {
		fmt.Println("GetAllUsers failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return c
	}
	c.JSON(http.StatusOK, users)
	return c
}
