package mirror

import (
	"errors"
	"fmt"
	"reflect"
)

func Fields(input any) []string {
	var fields []string
	value := reflect.ValueOf(input)

	for i := 0; i < value.NumField(); i++ {
		fields = append(fields, value.Type().Field(i).Name)
	}

	return fields
}

// Returns -1 if field is not found
func FieldIndex(input any, key string) int {
	fields := Fields(input)

	for i, f := range fields {
		if f == key {
			return i
		}
	}

	return -1
}

func Methods(s any) []string {
	var methods []string
	t := reflect.TypeOf(s)

	for i := 0; i < t.NumMethod(); i++ {
		methods = append(methods, t.Method(i).Name)
	}

	return methods
}

func SetField(input any, key string, value any) error {
	rv := reflect.ValueOf(input)
	if rv.Kind() != reflect.Pointer || rv.Elem().Kind() != reflect.Struct {
		return errors.New("input must be pointer to struct")
	}

	field := rv.Elem().FieldByName(key)
	if !field.IsValid() {
		return fmt.Errorf("not a field name: %s", key)
	}

	if !field.CanSet() {
		return fmt.Errorf("cannot set field %s", key)
	}

	field.Set(reflect.ValueOf(value))

	return nil
}

func GetField(input any, key string) (string, error) {
	value := reflect.ValueOf(input)
	if value.Kind() != reflect.Pointer || value.Elem().Kind() != reflect.Struct {
		return "", errors.New("input must be pointer to struct")
	}

	field := value.Elem().FieldByName(key)
	if !field.IsValid() {
		return "", fmt.Errorf("not a field name: %s", key)
	}

	return fmt.Sprintf("%s", field), nil
}
