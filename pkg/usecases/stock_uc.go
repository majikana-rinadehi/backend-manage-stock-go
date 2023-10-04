package usecases

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/entities"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces"
	"github.com/majikana-rinadehi/backend-manage-stock-go/pkg/interfaces/usecases"
)

type StockUsecase struct {
	r interfaces.StockRepository
}

func NewStockUsecase(r interfaces.StockRepository) usecases.StockUsecase {
	return &StockUsecase{
		r: r,
	}
}

func (uc StockUsecase) GetStocks(userId int) ([]*entities.Stock, error) {
	stocks, err := uc.r.FindByQuery(userId)
	if err != nil {
		fmt.Println("Failed GetAllStock; ", err)
		return nil, err
	}
	return stocks, nil
}

func (uc StockUsecase) CreateStock(stock *entities.Stock) (*entities.Stock, error) {
	stockSaved, err := uc.r.Save(stock)
	if err != nil {
		fmt.Println("Failed CreateStock;", err)
		return nil, err
	}
	return stockSaved, nil
}

func (uc StockUsecase) DeleteStock(stockId int) error {
	err := uc.r.DeleteById(stockId)
	if err != nil {
		fmt.Println("Failed DeleteStock;", err)
		return err
	}
	return nil
}

func (uc StockUsecase) UpdateStock(id int, stock *entities.Stock) (*entities.Stock, error) {

	count, err := uc.r.CountById(id)
	if err != nil {
		fmt.Println("Failed UpdateStock;", err)
		return nil, err
	}

	if count == 0 {
		fmt.Println("Update target not found;")
		return nil, errors.New("Not found: id = " + strconv.Itoa(id))
	}

	stockUpdated, err := uc.r.Update(id, stock)
	if err != nil {
		fmt.Println("Failed UpdateStock;", err)
		return nil, err
	}
	return stockUpdated, nil
}
