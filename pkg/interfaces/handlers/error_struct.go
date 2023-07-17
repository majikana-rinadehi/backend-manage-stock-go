package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/majikana-rinadehi/backend-manage-stock-go/util"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func (s ErrorResponse) String() string {
	return util.CustomStringer(s)
}

func BadRequest(c *gin.Context, vErr validation.ErrorObject) {
	c.JSON(http.StatusBadRequest, &Response[any]{
		Total:   0,
		Results: nil,
		Errors: []*ErrorResponse{
			{
				Message: vErr.Error(),
			},
		},
	})
}

func BadRequests(c *gin.Context, vErr validation.Errors) {
	errors := make([]*ErrorResponse, 0)
	for _, e := range vErr {
		errors = append(errors, &ErrorResponse{
			Message: e.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, &Response[any]{
		Total:   0,
		Results: nil,
		Errors:  SortErrorResponse(errors),
	})
}

func SortErrorResponse(errorList []*ErrorResponse) []*ErrorResponse {
	sorted := make([]*ErrorResponse, len(errorList))
	// 配列のディープコピー
	for i, err := range errorList {
		newErr := &ErrorResponse{
			Message: err.Message,
		}
		sorted[i] = newErr
	}
	// sort.Slice, sort.SliceStableではうまく並び替えられなかった
	sortByMessage(sorted)
	return sorted
}

// sortByMessageは、ErrorResponseの配列をMessageの辞書順(昇順)で並び替えます。
// バブルソートのアルゴリズムを使用しています。
func sortByMessage(list []*ErrorResponse) {
	length := len(list)
	if length <= 1 {
		return
	}

	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if list[j].Message < list[minIndex].Message {
				minIndex = j
			}
		}
		if minIndex != i {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}
