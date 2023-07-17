package handlers

import (
	"fmt"
	"reflect"
)

type Response[T any] struct {
	Total   int              `json:"total"`
	Results []T              `json:"results"`
	Errors  []*ErrorResponse `json:"errors"`
}

func (res Response[T]) String() string {
	output := "{\n"
	val := reflect.ValueOf(res)
	t := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldValue := field.Interface()
		tagName := t.Field(i).Tag.Get("json")
		output += fmt.Sprintf("\t%s: %v,\n",
			tagName, fieldValue)

	}

	output += "}"
	return output
}
