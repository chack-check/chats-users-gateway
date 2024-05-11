package chats

import (
	"fmt"
	"slices"

	"github.com/chack-check/chats-users-gateway/domain/users"
)

var (
	ErrChatNotFound = fmt.Errorf("Chat with this id doesn't exist")
)

func addUsersIds(userIds []int, newIds []int) []int {
	for _, id := range newIds {
		if slices.Contains(userIds, id) {
			continue
		}

		userIds = append(userIds, id)
	}

	return userIds
}

type GetConcreteChatHandler struct {
	chatsPort ChatsPort
	usersPort users.UsersPort
}

func (handler *GetConcreteChatHandler) Execute(id int, userId int) (*ChatWithUsers, error) {
	chat := handler.chatsPort.GetById(id)
	if chat == nil {
		return nil, ErrChatNotFound
	}

	var usersIds []int
	usersIds = addUsersIds(usersIds, chat.GetMembersIds())
	users := handler.usersPort.GetByIds(usersIds)
	return &ChatWithUsers{Chat: *chat, Users: users}, nil
}

type GetChatsByIdsHandler struct {
	chatsPort ChatsPort
	usersPort users.UsersPort
}

func (handler *GetChatsByIdsHandler) Execute(ids []int, userId int) ChatsWithUsers {
	chats := handler.chatsPort.GetByIds(ids)

	var usersIds []int
	for _, chat := range chats {
		usersIds = addUsersIds(usersIds, chat.GetMembersIds())
	}

	users := handler.usersPort.GetByIds(usersIds)
	return ChatsWithUsers{Chats: chats, Users: users}
}
