package util

import (
	"fmt"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	requiredErrMsgFormat    = "Param '%s' is required."
	invalidTypeErrMsgFormat = "Param '%s' must be a '%s'."
	maxLengthErrMsgFormat   = "Param '%s' length must be <= '%d'."
)

func RequiredErrMsg(field string) string {
	return fmt.Sprintf(requiredErrMsgFormat, field)
}

func InvalidTypeErrMsg(field, expectedType string) string {
	return fmt.Sprintf(invalidTypeErrMsgFormat, field, expectedType)
}

func MaxLengthErrMsg(field string, maxLength int) string {
	return fmt.Sprintf(maxLengthErrMsgFormat, field, maxLength)
}

func ValidateIntNotEmpty(fieldName string) validation.RuleFunc {
	return func(value interface{}) error {

		// 数値の場合
		v, okInt := value.(int)
		// FIXME: json上でのundefinedとzerovalueが現状区別できないから、0はリクエスト禁止www
		if !okInt || v == 0 {
			return validation.NewError("Required", RequiredErrMsg(fieldName))
		}

		return nil
	}
}

func ValidateStrNotEmpty(fieldName string) validation.RuleFunc {
	return func(value interface{}) error {

		// string型の場合
		_, ok := value.(string)
		if !ok {
			return validation.NewError("Required", RequiredErrMsg(fieldName))
		}

		if strings.TrimSpace(value.(string)) == "" {
			return validation.NewError("Required", RequiredErrMsg(fieldName))
		}
		return nil
	}
}

func ValidateYYYY_MM_DD(fieldName string) validation.RuleFunc {
	return func(value interface{}) error {

		dateStr, ok := value.(string)
		if !ok {
			return validation.NewError("Required", RequiredErrMsg(fieldName))
		}

		if strings.TrimSpace(dateStr) == "" {
			return validation.NewError("Required", RequiredErrMsg(fieldName))
		}

		// 日付の解析を試みる
		_, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return validation.NewError("InvalidDate", InvalidTypeErrMsg(fieldName, "YYYY-MM-DD"))
		}

		return nil
	}
}
