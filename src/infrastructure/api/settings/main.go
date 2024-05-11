package settings

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type SettingsSchema struct {
	APP_SECRET_KEY    string
	APP_ALLOW_ORIGINS string
	APP_PORT          int
}

func InitSettings() SettingsSchema {
	key := os.Getenv("APP_SECRET_KEY")
	if key == "" {
		panic(fmt.Errorf("you need to specify `APP_SECRET_KEY` environment variable"))
	}

	allowOrigins := os.Getenv("APP_ALLOW_ORIGINS")
	if key == "" {
		allowOrigins = "*"
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(errors.Join(fmt.Errorf("error parsing app port value"), err))
	}

	return SettingsSchema{
		APP_SECRET_KEY:    key,
		APP_ALLOW_ORIGINS: allowOrigins,
		APP_PORT:          portInt,
	}
}

var Settings SettingsSchema = InitSettings()
