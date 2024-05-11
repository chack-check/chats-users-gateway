package protousers

import (
	"fmt"
	"os"
	"strconv"
)

type SettingsSchema struct {
	APP_USERS_GRPC_HOST string
	APP_USERS_GRPC_PORT int
}

func InitSettings() SettingsSchema {
	host := os.Getenv("APP_USERS_GRPC_HOST")
	if host == "" {
		panic(fmt.Errorf("you need to specify `APP_USERS_GRPC_HOST` environment variable"))
	}

	usersPort := os.Getenv("APP_USERS_GRPC_PORT")
	if usersPort == "" {
		panic(fmt.Errorf("you need to specify `APP_USERS_GRPC_PORT` environment variable"))
	}
	usersPortInt, err := strconv.Atoi(usersPort)
	if err != nil {
		panic(err)
	}

	return SettingsSchema{
		APP_USERS_GRPC_HOST: host,
		APP_USERS_GRPC_PORT: usersPortInt,
	}
}

var Settings SettingsSchema = InitSettings()
