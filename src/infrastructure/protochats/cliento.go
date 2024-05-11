package protochats

import (
	"fmt"

	protobuf "github.com/chack-check/chats-users-gateway/infrastructure/protochats/chatsprotobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

var connection *grpc.ClientConn

func ChatsClientConnect() protobuf.ChatsClient {
	if connection == nil || connection.GetState() != connectivity.Ready {
		opts := grpc.WithTransportCredentials(insecure.NewCredentials())
		dsl := fmt.Sprintf("%s:%d", Settings.APP_CHATS_GRPC_HOST, Settings.APP_CHATS_GRPC_PORT)

		newConnection, err := grpc.Dial(dsl, opts)
		if err != nil {
			return nil
		}

		connection = newConnection
	}

	return protobuf.NewChatsClient(connection)
}
