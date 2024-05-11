package protochats

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type ChatsSettings struct {
	APP_CHATS_GRPC_HOST string
	APP_CHATS_GRPC_PORT int
}

func NewChatsSettings() ChatsSettings {
	host := os.Getenv("APP_CHATS_GRPC_HOST")
	if host == "" {
		panic(fmt.Errorf("you need to specify `APP_CHATS_GRPC_HOST` environment variable"))
	}

	port := os.Getenv("APP_CHATS_GRPC_PORT")
	if port == "" {
		panic(fmt.Errorf("you need to specify `APP_CHATS_GRPC_PORT` environment variable"))
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(errors.Join(fmt.Errorf("error parsing chats grpc port"), err))
	}

	return ChatsSettings{
		APP_CHATS_GRPC_HOST: host,
		APP_CHATS_GRPC_PORT: portInt,
	}
}

var Settings = NewChatsSettings()
