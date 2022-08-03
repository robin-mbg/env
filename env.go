package env

import (
	"fmt"
	"os"
	"strconv"
)

func SetString(key string, value string) error {
	return os.Setenv(key, value)
}

func SetInt(key string, value int) error {
	return SetString(key, strconv.Itoa(value))
}

func SetFloat64(key string, value float64) error {
	return SetString(key, fmt.Sprintf("%f", value))
}

func SetBool(key string, value bool) error {
	return SetString(key, fmt.Sprintf("%v", value))
}

func GetStringOrError(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %v is empty", key)
	}

	return value, nil
}

func GetStringOrDefault(key string, defaultValue string) string {
	value, err := GetStringOrError(key)
	if err != nil {
		return defaultValue
	}

	return value
}

func GetStringOrPanic(key string) string {
	value, err := GetStringOrError(key)
	if err != nil {
		panic(err)
	}

	return value
}

func GetBoolOrError(key string) (bool, error) {
	value := os.Getenv(key)
	if value == "" {
		return false, fmt.Errorf("environment variable %v is empty", key)
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("environment variable %v could not be parsed to boolean value", key)
	}

	return boolValue, nil
}

func GetBoolOrDefault(key string, defaultValue bool) bool {
	value, err := GetBoolOrError(key)
	if err != nil {
		return defaultValue
	}

	return value
}

func GetBoolOrPanic(key string) bool {
	value, err := GetBoolOrError(key)
	if err != nil {
		panic(err)
	}

	return value
}

func GetIntOrError(key string) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, fmt.Errorf("environment variable %v is empty", key)
	}

	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("environment variable %v could not be parsed to int value", key)
	}

	return int(intValue), nil
}

func GetIntOrDefault(key string, defaultValue int) int {
	value, err := GetIntOrError(key)
	if err != nil {
		return defaultValue
	}

	return value
}

func GetIntOrPanic(key string) int {
	value, err := GetIntOrError(key)
	if err != nil {
		panic(err)
	}

	return value
}

func GetFloat64OrError(key string) (float64, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, fmt.Errorf("environment variable %v is empty", key)
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("environment variable %v could not be parsed to int value", key)
	}

	return floatValue, nil
}

func GetFloat64OrDefault(key string, defaultValue float64) float64 {
	value, err := GetFloat64OrError(key)
	if err != nil {
		return defaultValue
	}

	return value
}

func GetFloat64OrPanic(key string) float64 {
	value, err := GetFloat64OrError(key)
	if err != nil {
		panic(err)
	}

	return value
}
