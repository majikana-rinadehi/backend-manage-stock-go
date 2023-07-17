package util

import (
	"fmt"
	"reflect"
)

func CustomStringer(s interface{}) string {
	output := "\n\t\t{\n"
	val := reflect.ValueOf(s)
	t := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldValue := field.Interface()
		tagName := t.Field(i).Tag.Get("json")
		output += fmt.Sprintf("\t\t\t%s: %v\n",
			tagName, fieldValue)
	}

	output += "\t\t},\n"

	return output
}
