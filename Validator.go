package folcon

import (
	"fmt"
	"reflect"
	"strings"
)

// ValidateRequired parses JSON and dynamically checks struct fields for 'binding:"required"' tag
func (c *Context) ValidateRequired(s interface{}) map[string]string {
	errors := make(map[string]string)

	// 1. Context-ல் உள்ள BindJSON மெதட்டை சரியாக அழைத்தல்
	if err := c.BindJSON(s); err != nil {
		errors["json_error"] = err.Error()
		return errors
	}

	// 2. ரிப்ளெக்சன் வேலிடேஷன் லாஜிக்
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

		// 'binding:"required"' டேக்கைச் சரிபார்த்தல்
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
