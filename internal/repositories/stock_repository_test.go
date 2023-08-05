package repositories

import (
	"reflect"
	"testing"

	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/adapters"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	testutil "github.com/majikana-rinadehi/backend-manage-stock-go/util/testUtil"
)

func TestStockRepository_FindAll(t *testing.T) {

	testutil.SetupFixtures()

	a := adapters.NewDatabaseAdapter()
	a.Connect(true)

	type fields struct {
		dbAdapter *adapters.DatabaseAdapter
	}
	tests := []struct {
		name       string
		fields     fields
		wantStocks []*entities.Stock
		wantErr    error
	}{
		// TODO: Add test cases.
		{
			name: "NORMAL",
			fields: fields{
				dbAdapter: a,
			},
			wantStocks: []*entities.Stock{
				{
					Id:         1,
					UserId:     1,
					CategoryId: 1,
					Name:       "きゅうり",
					Amount:     1,
					ExpireDate: "2023-01-01T09:00:00+09:00",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &StockRepository{
				dbAdapter: tt.fields.dbAdapter,
			}
			gotStocks, err := r.FindAll()
			if err != nil && err != tt.wantErr {
				t.Errorf("StockRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStocks, tt.wantStocks) {
				t.Errorf("StockRepository.FindAll() = %v, want %v", gotStocks, tt.wantStocks)
			}
		})
	}
}

func TestStockRepository_Save(t *testing.T) {

	testutil.SetupFixtures()

	a := adapters.NewDatabaseAdapter()
	a.Connect(true)

	type fields struct {
		dbAdapter *adapters.DatabaseAdapter
	}
	type args struct {
		stock *entities.Stock
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Stock
		wantErr error
	}{
		// TODO: Add test cases.
		// normal
		{
			name: "NORMAL",
			fields: fields{
				dbAdapter: a,
			},
			args: args{
				stock: &entities.Stock{
					Id:         999,
					UserId:     999,
					CategoryId: 999,
					Name:       "テスト",
					Amount:     1,
					ExpireDate: "2023-01-01T00:00:00+09:00",
				},
			},
			want: &entities.Stock{
				Id:         999,
				UserId:     999,
				CategoryId: 999,
				Name:       "テスト",
				Amount:     1,
				ExpireDate: "2023-01-01T00:00:00+09:00",
			},
			wantErr: nil,
		},
		// insert error
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &StockRepository{
				dbAdapter: tt.fields.dbAdapter,
			}
			got, err := r.Save(tt.args.stock)
			if err != nil && err != tt.wantErr {
				t.Errorf("StockRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StockRepository.Save() = %v, want %v", got, tt.want)
			}

			// テストデータがインサートされていることを確認
			// var insertedStock entities.Stock
			gormDb, _ := tt.fields.dbAdapter.GetDB()
			// db, _ := gormDb.DB()
			// rows := db.QueryRow("select * from stocks where id = ?", tt.args.stock.Id)
			// rows.Scan(&insertedStock)

			var insertedStocks []*entities.Stock
			gormDb.Find(&insertedStocks, tt.args.stock.Id)

			if !reflect.DeepEqual(tt.args.stock, insertedStocks[0]) {
				t.Errorf("StockRepository.Save() :Correct row sholud be inserted."+
					"want = %v, got = %v", tt.args.stock, insertedStocks[0])
			}

			t.Logf("insertedStocks: %v", insertedStocks)
		})
	}
}

func TestStockRepository_DeleteById(t *testing.T) {

	testutil.SetupFixtures()

	a := adapters.NewDatabaseAdapter()
	a.Connect(true)

	type fields struct {
		dbAdapter *adapters.DatabaseAdapter
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
		{
			"NORMAL",
			fields{
				dbAdapter: a,
			},
			args{
				id: 1,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &StockRepository{
				dbAdapter: tt.fields.dbAdapter,
			}
			if err := r.DeleteById(tt.args.id); err != nil && err != tt.wantErr {
				t.Errorf("StockRepository.DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
			// テストデータが削除されていることを確認
			var count int
			gormDb, _ := tt.fields.dbAdapter.GetDB()
			db, _ := gormDb.DB()
			db.QueryRow("select count(*) from stocks where id = ?", tt.args.id).Scan(&count)

			if !reflect.DeepEqual(count, 0) {
				t.Errorf("StockRepository.DeleteById() :Row sholud be deleted")
			}

			t.Logf("count: %d", count)

		})
	}
}
