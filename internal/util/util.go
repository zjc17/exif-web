package util

import "os"

func TrimByteArray(data []byte, length int) []byte {

	if len(data) <= length {
		return data
	} else {
		result := make([]byte, length)
		copy(result, data)
		return result
	}
}

func GetEnvOrDefault(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return defaultValue
	}
}
