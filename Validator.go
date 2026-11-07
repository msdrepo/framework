package folcon

import (
	"fmt"
	"reflect"
	"strings"
)

// validate required dynamically checks struct fields for 'binding:"required"' tag
func ValidateRequired(s interface{}) map[string]string {

	errors := make(map[string]string)
	val := reflect.ValueOf(s)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return errors
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)

		// 'binding:"required"'
		bindingTag := fieldType.Tag.Get("binding")
		if strings.Contains(bindingTag, "required") {

			jsonTag := fieldType.Tag.Get("json")
			fieldName := strings.Split(jsonTag, ",")[0]
			if fieldName == "" {
				fieldName = fieldType.Name
			}

			if fieldVal.Kind() == reflect.String && fieldVal.String() == "" {
				errors[fieldName] = fmt.Sprintf("%s is required", fieldName)
			}
		}
	}

	return errors
}
