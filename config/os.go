package config

import (
	"log"
	"os"
	"strconv"
)

func GetInt(key string, defaultValue int) int {
	i, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Printf("getting error from parsing config with key: %v and value: %v\n", key, i)
		return defaultValue
	}
	return i
}

func GetString(key string) string {
	return os.Getenv(key)
}

func GetBool(key string) bool {
	res := false
	if os.Getenv(key) == "true" || os.Getenv(key) == "t" {
		res = true
	}
	return res
}
