package getenv

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

// Get .env variable with value int
func GetInt(key string, defaultVal int) int {
	str := os.Getenv(key)
	if str == "" {
		return defaultVal
	}
	// convert string to int
	val, err := strconv.ParseInt(str, 10, 10)
	if err != nil {
		return defaultVal
	}
	return int(val)
}

// Get .env variable with value string
func GetString(key string, defaultVal string) string {
	str := os.Getenv(key)
	if str == "" {
		return defaultVal
	}
	return str
}

// get .env with value bool
func GetBool(key string,defaultVal bool) bool {
	str := os.Getenv(key)	
	if str == "" {
		return defaultVal
	}
	val, err := strconv.ParseBool(str)
	if err != nil {
		return defaultValue
	}

	return val
}