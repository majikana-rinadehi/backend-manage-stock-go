package adapters

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseAdapter struct {
	interfaces.Database
	db *gorm.DB
}

func NewDatabaseAdapter() *DatabaseAdapter {
	return &DatabaseAdapter{}
}

func (a *DatabaseAdapter) Connect(isTest bool) error {
	// データベースに接続する処理
	dsn := a.GetDSN(isTest)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db.Logger = db.Logger.LogMode(logger.Info)
	if err != nil {
		return err
	}
	a.db = db
	return nil
}

func (a *DatabaseAdapter) Disconnect() error {
	// データベースから切断する処理
	db, err := a.GetDB()
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	return sqlDB.Close()
}

func (a *DatabaseAdapter) GetDB() (*gorm.DB, error) {
	// GORMのDBインスタンスを返す処理
	if a.db == nil {
		return nil, fmt.Errorf("database connection not established")
	}
	return a.db, nil
}

func (a *DatabaseAdapter) GetDSN(isTest bool) string {
	env := os.Getenv("MANAGE_STOCK_ENV")

	switch env {
	case "production":
		env = "production"
	default:
		env = "development"
	}

	if isTest {
		os.Setenv("DBUSER", "root")
		os.Setenv("DBPASS", "root")
		// ↓↓↓ここはテスト用(DB名に「test」を含むこと)にしないと、
		// panic: testfixtures: database "manage_stock" does not appear to be a test database
		os.Setenv("DBNAME", "manage_stock_test")
		os.Setenv("DBHOST", "localhost")
		os.Setenv("DBPORT_TEST", "3307")
	} else {
		godotenv.Load(".env." + env)
	}

	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASS")
	DBNAME := os.Getenv("DBNAME")
	DBHOST := os.Getenv("DBHOST")
	var DBPORT string
	if isTest {
		DBPORT = os.Getenv("DBPORT_TEST")
	} else {
		DBPORT = os.Getenv("DBPORT")

	}

	CONNECT := DBUSER + ":" + DBPASS + "@tcp(" + DBHOST + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"
	fmt.Println(CONNECT)

	return CONNECT
}
