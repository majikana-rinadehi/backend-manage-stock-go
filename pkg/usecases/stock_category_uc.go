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
	rsc interfaces.StockCategoryRepository
	rs  interfaces.StockRepository
}

func NewStockCategoryUsecase(rsc interfaces.StockCategoryRepository, rs interfaces.StockRepository) usecases.StockCategoryUsecase {
	return &StockCategoryUsecase{
		rsc: rsc,
		rs:  rs,
	}
}

func (uc StockCategoryUsecase) GetStockCategories(userId int) ([]*entities.StockCategory, error) {
	stocks, err := uc.rsc.FindByQuery(userId)
	if err != nil {
		fmt.Println("Failed GetAllStockCategory; ", err)
		return nil, err
	}
	return stocks, nil
}

func (uc StockCategoryUsecase) CreateCategory(category *entities.StockCategory) (*entities.StockCategory, error) {
	category, err := uc.rsc.Save(category)
	if err != nil {
		fmt.Println("Failed CreateCategory; ", err)
		return nil, err
	}
	return category, nil
}

func (uc StockCategoryUsecase) DeleteCategory(categoryId int) error {
	err := uc.rsc.DeleteById(categoryId)
	if err != nil {
		fmt.Println("Failed DeleteCategory; ", err)
		return err
	}

	// カテゴリIDに紐づくStockをすべて削除する。
	err2 := uc.rs.DeleteByCategoryId(categoryId)
	if err2 != nil {
		fmt.Println("Failed DeleteCategory; ", err2)
		return err
	}

	return nil
}

func (uc StockCategoryUsecase) UpdateCategory(id int, category *entities.StockCategory) (*entities.StockCategory, error) {

	count, err := uc.rsc.CountById(id)
	if err != nil {
		fmt.Println("Failed UpdateCategory;", err)
		return nil, err
	}

	if count == 0 {
		fmt.Println("Update target not found;")
		return nil, errors.New("Not found: id = " + strconv.Itoa(id))
	}

	categoryUpdated, err := uc.rsc.Update(id, category)
	if err != nil {
		fmt.Println("Failed UpdateCategory;", err)
		return nil, err
	}
	return categoryUpdated, nil
}
