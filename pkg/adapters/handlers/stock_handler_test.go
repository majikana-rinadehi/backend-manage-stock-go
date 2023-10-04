package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/majikana-rinadehi/backend-manage-stock-go/docs"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/handlers"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/usecases"
	"github.com/majikana-rinadehi/backend-manage-stock-go/util"
	"github.com/stretchr/testify/mock"
)

type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) GetAllStocks() ([]*entities.Stock, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Stock), args.Error(1)
}

func (m *MockUsecase) CreateStock(stock *entities.Stock) (*entities.Stock, error) {
	args := m.Called()
	return args.Get(0).(*entities.Stock), args.Error(1)
}

func (m *MockUsecase) DeleteStock(stockId int) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUsecase) UpdateStock(stockId int, stock *entities.Stock) (*entities.Stock, error) {
	args := m.Called()
	return args.Get(0).(*entities.Stock), args.Error(1)
}

var mockStockList = []*entities.Stock{
	{
		Id:         1,
		UserId:     1,
		CategoryId: 1,
		Name:       "きゅうり",
		Amount:     1,
		ExpireDate: "2023-07-14",
	},
	{
		Id:         2,
		UserId:     2,
		CategoryId: 2,
		Name:       "酢",
		Amount:     2,
		ExpireDate: "2023-07-14",
	},
}

var mockStock = entities.Stock{
	Id:         1,
	UserId:     1,
	CategoryId: 1,
	Name:       "きゅうり",
	Amount:     1,
	ExpireDate: "2023-07-14",
}

func TestStockHandler_GetAllStocks(t *testing.T) {

	mockUsecase, mockUsecaseErr := new(MockUsecase), new(MockUsecase)

	// TODO: implement mock method
	mockUsecase.On("GetAllStocks").Return(mockStockList, nil)
	mockUsecaseErr.On("GetAllStocks").Return([]*entities.Stock{}, errors.New("db error"))
	type fields struct {
		stockUsecase usecases.StockUsecase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *gin.Context
		wantStatus int
		wantBody   handlers.Response[*entities.Stock]
	}{
		// TODO: Add test cases.
		{
			name: "200",
			fields: fields{
				stockUsecase: mockUsecase,
			},
			wantStatus: 200,
			wantBody: handlers.Response[*entities.Stock]{
				Total:   len(mockStockList),
				Results: mockStockList,
				Errors:  nil,
			},
		},
		{
			name: "500",
			fields: fields{
				stockUsecase: mockUsecaseErr,
			},
			wantStatus: 500,
			wantBody: handlers.Response[*entities.Stock]{
				Total:   0,
				Results: nil,
				Errors: []*handlers.ErrorResponse{
					{
						Message: "GetAllStocks failed",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &StockHandler{
				stockUsecase: tt.fields.stockUsecase,
			}

			// Setup gin context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			got := h.GetStocks(c)

			// Assert status code
			if !reflect.DeepEqual(got.Writer.Status(), tt.wantStatus) {
				t.Errorf("StockHandler.GetAllStocks() = %v, wantStatus %v", got.Writer.Status(), tt.wantStatus)
			}

			// Assert response body
			var gotBody handlers.Response[*entities.Stock]
			json.Unmarshal(w.Body.Bytes(), &gotBody)
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("StockHandler.GetAllStocks() = %v, wantBody %v", gotBody, tt.wantBody)
			}

			t.Logf("TestStockHandler_GetAllStocks: case「%v」", tt.name)
			t.Logf("Status: %v", got.Writer.Status())
			t.Logf("Response: %v", fmt.Sprintf("%+v", gotBody))
		})
	}
}

func TestStockHandler_CreateStock(t *testing.T) {

	mockUsecase, mockUsecaseErr := new(MockUsecase), new(MockUsecase)

	// TODO: implement mock method
	mockUsecase.On("CreateStock").Return(&mockStock, nil)
	mockUsecaseErr.On("CreateStock").Return(&entities.Stock{}, errors.New("db error"))
	type fields struct {
		stockUsecase usecases.StockUsecase
	}
	tests := []struct {
		name       string
		fields     fields
		args       *entities.Stock
		want       *gin.Context
		wantStatus int
		wantBody   handlers.Response[*entities.Stock]
	}{
		// TODO: Add test cases.
		{
			name: "200",
			fields: fields{
				stockUsecase: mockUsecase,
			},
			args:       &mockStock,
			wantStatus: 200,
			wantBody: handlers.Response[*entities.Stock]{
				Total:   1,
				Results: []*entities.Stock{&mockStock},
				Errors:  nil,
			},
		},
		{
			name: "400_empty",
			fields: fields{
				stockUsecase: mockUsecase,
			},
			args:       &entities.Stock{},
			wantStatus: 400,
			wantBody: handlers.Response[*entities.Stock]{
				Total:   0,
				Results: nil,
				Errors: handlers.SortErrorResponse([]*handlers.ErrorResponse{
					{
						Message: util.RequiredErrMsg("categoryId"),
					},
					{
						Message: util.RequiredErrMsg("userId"),
					},
					{
						Message: util.RequiredErrMsg("amount"),
					},
					{
						Message: util.RequiredErrMsg("name"),
					},
					{
						Message: util.RequiredErrMsg("expireDate"),
					},
				}),
			},
		},
		{
			name: "400_invalidType",
			fields: fields{
				stockUsecase: mockUsecase,
			},
			args: &entities.Stock{
				Id:         1,
				UserId:     1,
				CategoryId: 1,
				// 256 length
				Name:       "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUV",
				Amount:     1,
				ExpireDate: "2023-99-99",
			},
			wantStatus: 400,
			wantBody: handlers.Response[*entities.Stock]{
				Total:   0,
				Results: nil,
				Errors: handlers.SortErrorResponse([]*handlers.ErrorResponse{
					{
						Message: util.MaxLengthErrMsg("name", 255),
					},
					{
						Message: util.InvalidTypeErrMsg("expireDate", "YYYY-MM-DD"),
					},
				}),
			},
		},
		{
			name: "500",
			fields: fields{
				stockUsecase: mockUsecaseErr,
			},
			args:       &mockStock,
			wantStatus: 500,
			wantBody: handlers.Response[*entities.Stock]{
				Total:   0,
				Results: nil,
				Errors: []*handlers.ErrorResponse{
					{
						Message: "CreateStock failed",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &StockHandler{
				stockUsecase: tt.fields.stockUsecase,
			}

			// Setup gin context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// リクエストボディを追加する
			jsonValue, _ := json.Marshal(tt.args)
			reqBody := bytes.NewBuffer(jsonValue)
			req, _ := http.NewRequest(
				"POST",
				"",
				reqBody,
			)
			c.Request = req

			got := h.CreateStock(c)

			// Assert status code
			if !reflect.DeepEqual(got.Writer.Status(), tt.wantStatus) {
				t.Errorf("StockHandler.CreateStock() = %v, wantStatus %v", got.Writer.Status(), tt.wantStatus)
			}

			// Assert response body
			var gotBody handlers.Response[*entities.Stock]
			json.Unmarshal(w.Body.Bytes(), &gotBody)
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("StockHandler.CreateStock() = %v, wantBody %v", gotBody, tt.wantBody)
			}

			t.Logf("TestStockHandler_CreateStock: case「%v」", tt.name)
			t.Logf("Status: %v", got.Writer.Status())
			t.Logf("Response: %+v", fmt.Sprintf("%+v", gotBody))
		})
	}
}

func TestStockHandler_DeleteStock(t *testing.T) {

	mockUsecase, mockUsecaseErr := new(MockUsecase), new(MockUsecase)

	// TODO: implement mock method
	mockUsecase.On("DeleteStock").Return(nil)
	mockUsecaseErr.On("DeleteStock").Return(errors.New("db error"))
	type fields struct {
		stockUsecase usecases.StockUsecase
	}
	tests := []struct {
		name       string
		fields     fields
		args       string
		want       *gin.Context
		wantStatus int
		wantBody   handlers.Response[any]
	}{
		// TODO: Add test cases.
		{
			name: "204",
			fields: fields{
				stockUsecase: mockUsecase,
			},
			args:       "1",
			wantStatus: 204,
			wantBody: handlers.Response[any]{
				Total:   0,
				Results: nil,
				Errors:  nil,
			},
		},
		{
			name: "500",
			fields: fields{
				stockUsecase: mockUsecaseErr,
			},
			args:       "1",
			wantStatus: 500,
			wantBody: handlers.Response[any]{
				Total:   0,
				Results: nil,
				Errors: []*handlers.ErrorResponse{
					{
						Message: "DeleteStock failed",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &StockHandler{
				stockUsecase: tt.fields.stockUsecase,
			}

			// Setup gin context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// パスパラメータを設定する
			c.Params = []gin.Param{
				{
					Key:   "id",
					Value: tt.args,
				},
			}
			got := h.DeleteStock(c)

			// Assert status code
			if !reflect.DeepEqual(got.Writer.Status(), tt.wantStatus) {
				t.Errorf("StockHandler.DeleteStock() = %v, wantStatus %v", got.Writer.Status(), tt.wantStatus)
			}

			// Assert response body
			var gotBody handlers.Response[any]
			json.Unmarshal(w.Body.Bytes(), &gotBody)
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("StockHandler.DeleteStock() = %v, wantBody %v", gotBody, tt.wantBody)
			}

			t.Logf("TestStockHandler_DeleteStock: case「%v」", tt.name)
			t.Logf("Status: %v", got.Writer.Status())
			t.Logf("Response: %v", fmt.Sprintf("%+v", gotBody))
		})
	}
}
