// アプリケーションのエントリポイントとなります
package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/majikana-rinadehi/backend-manage-stock-go/docs"
	"github.com/majikana-rinadehi/backend-manage-stock-go/internal/repositories"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/adapters"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/usecases"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// アダプターの取得、DB接続処理
	a := adapters.NewDatabaseAdapter()
	connErr := a.Connect(false)
	if connErr != nil {
		log.Fatal()
	}

	// リポジトリのセットアップ
	stockRepository := repositories.NewStockRepository(a)
	stockCategoryRepository := repositories.NewStockCategoryRepository(a)
	userRepository := repositories.NewUserRepository(a)

	// ユースケースのセットアップ
	stockUsecase := usecases.NewStockUsecase(stockRepository)
	stockCategoryUsecase := usecases.NewStockCategoryUsecase(stockCategoryRepository)
	userUsecase := usecases.NewUserUsecase(userRepository)

	// HTTPアダプターの作成
	httpAdapter := adapters.NewHTTPAdapter(stockUsecase, stockCategoryUsecase, userUsecase)

	router := gin.Default()
	// CORS設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))

	// ルーティングのセットアップ
	httpAdapter.SetupRoutes(router)

	// http://localhost:8080/swagger/index.html にswagger UI を表示する
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// サーバの起動
	log.Fatal(router.Run())
}
