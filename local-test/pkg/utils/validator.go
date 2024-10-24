package utils

import "fmt"

func ValidateField(objectName, fieldName string, fieldValue interface{},) error {
	switch v := fieldValue.(type) {
	case string:
		if v == "" {
			return fmt.Errorf("%s %s is required", objectName, fieldName)
		}
	case int:
		if v == 0 {
			return fmt.Errorf("%s %s is required", objectName, fieldName)
		}
	default:
		if fieldValue == nil {
			return fmt.Errorf("%s %s is required", objectName, fieldName)
		}
	}
	return nil
}