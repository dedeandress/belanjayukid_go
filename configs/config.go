package configs

import (
	"fmt"
	"os"
)

func GetConfigRequired(key string) (value string, err error) {
	value = GetConfig(key)
	if len(value) < 1 {
		return "", fmt.Errorf("Cannot find config for key %s\n", key)
	}

	return value, nil
}

func GetConfig(key string) (value string) {
	return os.Getenv(key)
}
