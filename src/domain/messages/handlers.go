package messages

import (
	"slices"

	"github.com/chack-check/chats-users-gateway/domain/chats"
	"github.com/chack-check/chats-users-gateway/domain/users"
)

func getUsersIdsForMessage(message Message, chat chats.Chat) []int {
	var usersIds []int

	for _, user := range chat.GetMembersIds() {
		if slices.Contains(usersIds, user) {
			continue
		}

		usersIds = append(usersIds, user)
	}

	return usersIds
}

type GetChatMessagesHandler struct {
	messagesPort MessagesPort
	chatsPort    chats.ChatsPort
	usersPort    users.UsersPort
}

func (handler *GetChatMessagesHandler) Execute(chatId, userId, offset, limit int) PaginatedMessagesWithUsers {
	messages := handler.messagesPort.GetByChatId(chatId, offset, limit)
	var usersIds []int
	for _, message := range messages.Messages {
		chat := handler.chatsPort.GetById(message.GetChatId())
		if chat == nil {
			continue
		}

		messageUsersIds := getUsersIdsForMessage(message, *chat)
		usersIds = append(usersIds, messageUsersIds...)
	}

	users := handler.usersPort.GetByIds(usersIds)
	return PaginatedMessagesWithUsers{
		Limit:    messages.Limit,
		Offset:   messages.Offset,
		Total:    messages.Total,
		Messages: messages.Messages,
		Users:    users,
	}
}

func NewGetChatMessagesHandler(messagesPort MessagesPort, chatsPort chats.ChatsPort, usersPort users.UsersPort) GetChatMessagesHandler {
	return GetChatMessagesHandler{
		messagesPort: messagesPort,
		chatsPort:    chatsPort,
		usersPort:    usersPort,
	}
}
