// package env provides utility methods for accessing environment variables with error, panic, and default value handling.
package env

import (
	"fmt"
	"os"
	"strconv"
)

// Sets an environment variable to a string value
func SetString(key string, value string) error {
	return os.Setenv(key, value)
}


// Sets an environment variable to an int value
func SetInt(key string, value int) error {
	return SetString(key, strconv.Itoa(value))
}

// Sets an environment variable to a float64 value
func SetFloat64(key string, value float64) error {
	return SetString(key, fmt.Sprintf("%f", value))
}

// Sets an environment variable to a bool value
func SetBool(key string, value bool) error {
	return SetString(key, fmt.Sprintf("%v", value))
}

// Gets a string value from an environment variable, returns an error if that variable is empty
func GetStringOrError(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %v is empty", key)
	}

	return value, nil
}

// Gets a string value from an environment variable, returns a default if that variable is empty
func GetStringOrDefault(key string, defaultValue string) string {
	value, err := GetStringOrError(key)
	if err != nil {
		return defaultValue
	}

	return value
}

// Gets a string value from an environment variable, panics if that variable is empty
func GetStringOrPanic(key string) string {
	value, err := GetStringOrError(key)
	if err != nil {
		panic(err)
	}

	return value
}

// Gets a bool value from an environment variable, returns an error if that variable is empty
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

// Gets a bool value from an environment variable, returns a default value if that variable is empty
func GetBoolOrDefault(key string, defaultValue bool) bool {
	value, err := GetBoolOrError(key)
	if err != nil {
		return defaultValue
	}

	return value
}

// Gets a bool value from an environment variable, panics if that variable is empty
func GetBoolOrPanic(key string) bool {
	value, err := GetBoolOrError(key)
	if err != nil {
		panic(err)
	}

	return value
}

// Gets an int value from an environment variable, returns an error if that variable is empty
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

// Gets an int value from an environment variable, returns a default value if that variable is empty
func GetIntOrDefault(key string, defaultValue int) int {
	value, err := GetIntOrError(key)
	if err != nil {
		return defaultValue
	}

	return value
}

// Gets an int value from an environment variable, panics if that variable is empty
func GetIntOrPanic(key string) int {
	value, err := GetIntOrError(key)
	if err != nil {
		panic(err)
	}

	return value
}

// Gets a float64 value from an environment variable, returns an error if that variable is empty
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

// Gets a float64 value from an environment variable, returns a default value if that variable is empty
func GetFloat64OrDefault(key string, defaultValue float64) float64 {
	value, err := GetFloat64OrError(key)
	if err != nil {
		return defaultValue
	}

	return value
}

// Gets a float64 value from an environment variable, panics if that variable is empty
func GetFloat64OrPanic(key string) float64 {
	value, err := GetFloat64OrError(key)
	if err != nil {
		panic(err)
	}

	return value
}
