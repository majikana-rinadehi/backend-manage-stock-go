package adapters

import (
	"github.com/gin-gonic/gin"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/adapters/handlers"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/usecases"
)

type HTTPAdapter struct {
	stockUc         usecases.StockUsecase
	stockCategoryUc usecases.StockCategoryUsecase
	userUc          usecases.UserUsecase
}

func NewHTTPAdapter(stockUc usecases.StockUsecase, stockCategoryUc usecases.StockCategoryUsecase, userUc usecases.UserUsecase) *HTTPAdapter {
	return &HTTPAdapter{
		stockUc:         stockUc,
		stockCategoryUc: stockCategoryUc,
		userUc:          userUc,
	}
}

func (a *HTTPAdapter) SetupRoutes(router *gin.Engine) {

	stockHandler := handlers.NewStockHandler(a.stockUc)
	stockCategoryHandler := handlers.NewStockCategoryHandler(a.stockCategoryUc)
	userHandler := handlers.NewUserHandler(a.userUc)
	authHandler := handlers.NewAuthHandler(a.userUc)

	// router.GET() の第二引数のシグネチャに合わせるため、ハンドラー関数をラップします。
	router.GET("/stocks", func(c *gin.Context) {
		stockHandler.GetStocks(c)
	})
	router.POST("/stocks", func(c *gin.Context) {
		stockHandler.CreateStock(c)
	})
	router.DELETE("/stocks/:id", func(c *gin.Context) {
		stockHandler.DeleteStock(c)
	})
	router.PUT("/stocks/:id", func(c *gin.Context) {
		stockHandler.UpdateStock(c)
	})
	router.GET("/categories", func(c *gin.Context) {
		stockCategoryHandler.GetCategories(c)
	})
	router.POST("/categories", func(c *gin.Context) {
		stockCategoryHandler.CreateCategory(c)
	})
	router.DELETE("/categories/:id", func(c *gin.Context) {
		stockCategoryHandler.DeleteCategory(c)
	})
	router.PUT("/categories/:id", func(c *gin.Context) {
		stockCategoryHandler.UpdateCategory(c)
	})
	router.GET("/user-auths", func(c *gin.Context) {
		userHandler.GetUserAuths(c)
	})
	router.POST("/auth/signin", func(c *gin.Context) {
		authHandler.SignInWithProvider(c)
	})
	router.PUT("/users/:id", func(c *gin.Context) {
		userHandler.PutUser(c)
	})
}
