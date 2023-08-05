package testutil

import (
	"fmt"
	"github.com/go-testfixtures/testfixtures/v3"
	"path"
	"runtime"

	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/adapters"
)

const (
	fixturesDirRelativePathFormat = "%s/../../internal/repositories/fixtures"
)

func SetupFixtures() {

	a := new(adapters.DatabaseAdapter)
	a.Connect(true)
	db, _ := a.GetDB()
	sqlDb, _ := db.DB()

	_, pwd, _, _ := runtime.Caller(0)
	dir := fmt.Sprintf(fixturesDirRelativePathFormat, path.Dir(pwd))

	fmt.Println("dir:", dir)

	fixtures, err := testfixtures.New(
		testfixtures.Database(sqlDb),  // You database connection
		testfixtures.Dialect("mysql"), // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory(dir),   // The directory containing the YAML files
	)
	if err != nil {
		panic(err)
	}

	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}
