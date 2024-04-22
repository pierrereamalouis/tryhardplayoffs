package utils

import (
	"errors"
	"net/url"
	"reflect"
)

func SetQueryParams(params any, vals *url.Values) error {

	paramsType := reflect.TypeOf(params)
	paramsValues := reflect.ValueOf(params)

	// Check if the struct has at least one field
	if paramsType.NumField() < 1 {
		return errors.New("struct must have at least one field")
	}

	for i := 0; i < paramsType.NumField(); i++ {
		// Get the field
		field := paramsType.Field(i)

		// Get the value of the custom struct tag
		urlTag := field.Tag.Get("url")

		// Get the value of the field
		fieldValue := paramsValues.Field(i)

		// Skip if field value is nil
		if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
			continue
		}

		// Check if the field is not of type string
		if fieldValue.Kind() != reflect.String {
			return errors.New("all fields must be of type string")
		}

		vals.Set(urlTag, fieldValue.String())
	}

	return nil
}
