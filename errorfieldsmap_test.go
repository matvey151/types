package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewValidation(t *testing.T) {
	result := NewValidation()
	if result == nil {
		t.Fatal("Получен nil")
	}
}

func TestAddError(t *testing.T) {
	expectedErr := Validation{
		"key1": []string{"value1"},
		"key2": []string{"value2", "value3"},
	}
	err := make(Validation)
	err.AddError("key1", "value1")
	err.AddError("key2", "value2")
	err.AddError("key2", "value3")
	if !reflect.DeepEqual(expectedErr, err) {
		t.Fatalf("Структуры не равны. Ожидалось: \n%v\nПолучено: \n%v\n",
			expectedErr, err)
	}
}

func TestAddErrors(t *testing.T) {
	to := Validation{
		"key1": []string{"value1"},
		"key2": []string{"value2", "value3"},
	}
	from := Validation{
		"key1": []string{"value4"},
		"key3": []string{"value5"},
		"key4": []string{},
	}
	expectedResult := Validation{
		"key1": []string{"value1", "value4"},
		"key2": []string{"value2", "value3"},
		"key3": []string{"value5"},
		"key4": []string{},
	}
	to.AddErrors(from)
	if !reflect.DeepEqual(to, expectedResult) {
		t.Fatalf("Структуры не равны. Ожидалось: \n%v\nПолучено: \n%v\n",
			expectedResult, to)
	}
}

func TestHasErrors(t *testing.T) {
	hasErrors := Validation{
		"key0": []string{},
		"key1": []string{"value1"},
		"key2": []string{"value2", "value3"},
	}
	if !hasErrors.HasErrors() {
		t.Fatal("Должны были быть ошибки")
	}

	hasNoErrors := Validation{
		"key0": []string{},
	}
	if hasNoErrors.HasErrors() {
		t.Fatal("Не должно было быть ошибок")
	}
}

func TestString(t *testing.T) {
	hasErrors := Validation{
		"key0": []string{},
		"key1": []string{"value1"},
		"key2": []string{"value2", "value3"},
	}
	expectedString := "key1: value1\nkey2: value2\nkey2: value3"
	receivedString := fmt.Sprintf("%s", hasErrors)
	if expectedString != receivedString {
		t.Fatalf("Ожидалась строка: \n%s\nПолучено: \n%s\n", expectedString, receivedString)
	}

	hasNoErrors := Validation{
		"key0": []string{},
	}
	expectedString = ""
	receivedString = fmt.Sprintf("%s", hasNoErrors)
	if expectedString != receivedString {
		t.Fatalf("Ожидалась строка: \n%s\nПолучено: \n%s\n", expectedString, receivedString)
	}
}
