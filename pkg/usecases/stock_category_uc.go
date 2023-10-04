package usecases

import (
	"errors"
	"fmt"
	"strconv"

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

func (uc StockCategoryUsecase) GetStockCategories(userId int) ([]*entities.StockCategory, error) {
	stocks, err := uc.r.FindByQuery(userId)
	if err != nil {
		fmt.Println("Failed GetAllStockCategory; ", err)
		return nil, err
	}
	return stocks, nil
}

func (uc StockCategoryUsecase) CreateCategory(category *entities.StockCategory) (*entities.StockCategory, error) {
	category, err := uc.r.Save(category)
	if err != nil {
		fmt.Println("Failed CreateCategory; ", err)
		return nil, err
	}
	return category, nil
}

func (uc StockCategoryUsecase) DeleteCategory(categoryId int) error {
	err := uc.r.DeleteById(categoryId)
	if err != nil {
		fmt.Println("Failed CreateCategory; ", err)
		return err
	}
	return nil
}

func (uc StockCategoryUsecase) UpdateCategory(id int, category *entities.StockCategory) (*entities.StockCategory, error) {

	count, err := uc.r.CountById(id)
	if err != nil {
		fmt.Println("Failed UpdateCategory;", err)
		return nil, err
	}

	if count == 0 {
		fmt.Println("Update target not found;")
		return nil, errors.New("Not found: id = " + strconv.Itoa(id))
	}

	categoryUpdated, err := uc.r.Update(id, category)
	if err != nil {
		fmt.Println("Failed UpdateCategory;", err)
		return nil, err
	}
	return categoryUpdated, nil
}
