package env

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jgroeneveld/trial/assert"
)

// ------------ STRING values --------------------
func TestGetStringOrError_Then_string_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := "TEST_KEY_VALUE"
	SetString(testKeyName, testKeyValue)

	// when
	value, err := GetStringOrError(testKeyName)

	// then
	assert.MustBeNil(t, err)
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetStringOrError_When_empty_Then_error_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()

	// when
	value, err := GetStringOrError(testKeyName)

	// then
	assert.MustNotBeNil(t, err)
	assert.MustBeEqual(t, "", value)
}

func TestGetStringOrDefault_Then_string_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := "TEST_KEY_VALUE"
	SetString(testKeyName, testKeyValue)

	// when
	value := GetStringOrDefault(testKeyName, "UNUSED_DEFAULT")

	// then
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetStringOrDefault_When_empty_Then_default_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	defaultValue := "DEFAULT"

	// when
	value := GetStringOrDefault(testKeyName, defaultValue)

	// then
	assert.MustBeEqual(t, defaultValue, value)
}

func TestGetStringOrPanic_Then_string_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := "TEST_KEY_VALUE"
	SetString(testKeyName, testKeyValue)

	// when
	value := GetStringOrPanic(testKeyName)

	// then
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetStringOrPanic_When_empty_Then_panic(t *testing.T) {
	// given
	testKeyName := getTestKey()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected: panic, actual: no panic")
		}
	}()

	// when
	GetStringOrPanic(testKeyName)
}

// ------------ INT values --------------------
func TestGetIntOrError_Then_int_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := 5
	SetInt(testKeyName, testKeyValue)

	// when
	value, err := GetIntOrError(testKeyName)

	// then
	assert.MustBeNil(t, err)
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetIntOrError_When_empty_Then_error_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()

	// when
	value, err := GetIntOrError(testKeyName)

	// then
	assert.MustNotBeNil(t, err)
	assert.MustBeEqual(t, 0, value)
}

func TestGetIntOrError_When_of_wrong_type_Then_error_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := "STRING"
	SetString(testKeyName, testKeyValue)

	// when
	value, err := GetIntOrError(testKeyName)

	// then
	assert.MustNotBeNil(t, err)
	assert.MustBeEqual(t, 0, value)
}

func TestGetIntOrDefault_Then_int_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := 42
	SetInt(testKeyName, testKeyValue)

	// when
	value := GetIntOrDefault(testKeyName, 10000)

	// then
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetIntOrDefault_When_empty_Then_default_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	defaultValue := 42

	// when
	value := GetIntOrDefault(testKeyName, defaultValue)

	// then
	assert.MustBeEqual(t, defaultValue, value)
}

func TestGetIntOrPanic_Then_int_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := 42
	SetInt(testKeyName, testKeyValue)

	// when
	value := GetIntOrPanic(testKeyName)

	// then
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetIntOrPanic_When_empty_Then_panic(t *testing.T) {
	// given
	testKeyName := getTestKey()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected: panic, actual: no panic")
		}
	}()

	// when
	GetIntOrPanic(testKeyName)
}

// ------------ BOOL values --------------------
func TestGetBoolOrError_Then_bool_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := true
	SetBool(testKeyName, testKeyValue)

	// when
	value, err := GetBoolOrError(testKeyName)

	// then
	assert.MustBeNil(t, err)
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetBoolOrError_When_empty_Then_error_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()

	// when
	value, err := GetBoolOrError(testKeyName)

	// then
	assert.MustNotBeNil(t, err)
	assert.MustBeEqual(t, false, value)
}

func TestGetBoolOrError_When_of_wrong_type_Then_error_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := "STRING"
	SetString(testKeyName, testKeyValue)

	// when
	value, err := GetBoolOrError(testKeyName)

	// then
	assert.MustNotBeNil(t, err)
	assert.MustBeEqual(t, false, value)
}

func TestGetBoolOrDefault_Then_bool_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := true
	SetBool(testKeyName, testKeyValue)

	// when
	value := GetBoolOrDefault(testKeyName, false)

	// then
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetBoolOrDefault_When_empty_Then_default_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	defaultValue := true

	// when
	value := GetBoolOrDefault(testKeyName, defaultValue)

	// then
	assert.MustBeEqual(t, defaultValue, value)
}

func TestGetBoolOrPanic_Then_bool_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := true
	SetBool(testKeyName, testKeyValue)

	// when
	value := GetBoolOrPanic(testKeyName)

	// then
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetBoolOrPanic_When_empty_Then_panic(t *testing.T) {
	// given
	testKeyName := getTestKey()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected: panic, actual: no panic")
		}
	}()

	// when
	GetBoolOrPanic(testKeyName)
}

// ------------ FLOAT64 values --------------------
func TestGetFloat64OrError_Then_string_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := 12.345
	SetFloat64(testKeyName, testKeyValue)

	// when
	value, err := GetFloat64OrError(testKeyName)

	// then
	assert.MustBeNil(t, err)
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetFloat64OrError_When_empty_Then_error_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()

	// when
	value, err := GetFloat64OrError(testKeyName)

	// then
	assert.MustNotBeNil(t, err)
	assert.MustBeEqual(t, 0.0, value)
}

func TestGetFloat64OrError_When_of_wrong_type_Then_error_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := "STRING"
	SetString(testKeyName, testKeyValue)

	// when
	value, err := GetFloat64OrError(testKeyName)

	// then
	assert.MustNotBeNil(t, err)
	assert.MustBeEqual(t, 0.0, value)
}

func TestGetFloat64OrDefault_Then_float64_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := 1.2345
	SetFloat64(testKeyName, testKeyValue)

	// when
	value := GetFloat64OrDefault(testKeyName, 0.000)

	// then
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetFloat64OrDefault_When_empty_Then_default_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	defaultValue := 1.23456

	// when
	value := GetFloat64OrDefault(testKeyName, defaultValue)

	// then
	assert.MustBeEqual(t, defaultValue, value)
}

func TestGetFloat64OrPanic_Then_float64_is_returned(t *testing.T) {
	// given
	testKeyName := getTestKey()
	testKeyValue := 1.2345
	SetFloat64(testKeyName, testKeyValue)

	// when
	value := GetFloat64OrPanic(testKeyName)

	// then
	assert.MustBeEqual(t, testKeyValue, value)
}

func TestGetFloat64OrPanic_When_empty_Then_panic(t *testing.T) {
	// given
	testKeyName := getTestKey()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected: panic, actual: no panic")
		}
	}()

	// when
	GetFloat64OrPanic(testKeyName)
}

func getTestKey() string {
	return uuid.NewString()
}
