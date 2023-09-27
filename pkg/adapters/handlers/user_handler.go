package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	_ "github.com/majikana-rinadehi/backend-manage-stock-go/docs"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/handlers"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/usecases"
	"github.com/majikana-rinadehi/backend-manage-stock-go/util"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(uc usecases.UserUsecase) handlers.UserHandler {
	return &UserHandler{
		userUsecase: uc,
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

	users, err := h.userUsecase.GetUserAuths(userId, authProvider, uid)
	if err != nil {
		fmt.Println("GetAllUsers failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return c
	}
	c.JSON(http.StatusOK, users)
	return c
}

// PutUser
// @Summary userIdに対応するユーザの情報を更新する
// @Tags User
// @Produce json
// @Param id path string false "id"
// @Param body body entities.User false "User"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /users/{id} [put]
func (h *UserHandler) PutUser(c *gin.Context) *gin.Context {
	id, _ := strconv.Atoi(c.Param("id"))

	var newUser entities.User

	if bindErr := c.BindJSON(&newUser); bindErr != nil {
		c.JSON(http.StatusBadRequest, &handlers.Response[any]{
			Total:   0,
			Results: nil,
			Errors: []*handlers.ErrorResponse{
				{
					Message: bindErr.Error(),
				},
			},
		})
		return c
	}

	vErr := validateUser(newUser)

	if vErr != nil {
		handlers.BadRequests(c, vErr.(validation.Errors))
		return c
	}

	err := h.userUsecase.UpdateUser(id, &newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &handlers.Response[any]{
			Total:   0,
			Results: nil,
			Errors: []*handlers.ErrorResponse{
				{
					Message: "UpdateUser failed",
				},
			},
		})
		return c
	}

	c.Status(http.StatusNoContent)
	return c
}

// validateUser validates User data in request body
func validateUser(user entities.User) error {
	// リクエストバリデーションチェック
	vErr := validation.ValidateStruct(&user,
		validation.Field(&user.Name,
			validation.By(util.ValidateStrNotEmpty("name")),
		),
		validation.Field(&user.MailAddress,
			is.Email,
			validation.By(util.ValidateStrNotEmpty("mailAddress")),
		),
	)

	return vErr
}
