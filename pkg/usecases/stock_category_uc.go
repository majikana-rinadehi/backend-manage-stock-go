package usecases

import (
	"fmt"

	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/usecases"
)

type StockCategoryUsecase struct {
	r interfaces.StockCategoryRepository
}

func NewStockCategoryUsecase(r interfaces.StockCategoryRepository) usecases.StockCategoryUsecase {
	return &StockCategoryUsecase{
		r: r,
	}
}

func (uc StockCategoryUsecase) GetAllStockCategories() ([]*entities.StockCategory, error) {
	stocks, err := uc.r.FindAll()
	if err != nil {
		fmt.Println("Failed GetAllStockCategory; ", err)
		return nil, err
	}
	return stocks, nil
}
