package chats

import (
	"github.com/chack-check/chats-users-gateway/domain/files"
	"github.com/chack-check/chats-users-gateway/domain/users"
)

type ChatWithUsers struct {
	Chat  Chat
	Users []users.User
}

type ChatsWithUsers struct {
	Chats []Chat
	Users []users.User
}

type Chat struct {
	id         int
	avatar     *files.SavedFile
	title      string
	type_      string
	membersIds []int
	isArchived bool
	ownerId    int
	adminsIds  []int
}

func (c *Chat) GetId() int {
	return c.id
}

func (c *Chat) GetAvatar() *files.SavedFile {
	return c.avatar
}

func (c *Chat) GetTitle() string {
	return c.title
}

func (c *Chat) GetType() string {
	return c.type_
}

func (c *Chat) GetMembersIds() []int {
	return c.membersIds
}

func (c *Chat) GetIsArchived() bool {
	return c.isArchived
}

func (c *Chat) GetOwnerId() int {
	return c.ownerId
}

func (c *Chat) GetAdminsIds() []int {
	return c.adminsIds
}

func NewChat(
	id int,
	avatar *files.SavedFile,
	title string,
	type_ string,
	membersIds []int,
	isArchived bool,
	ownerId int,
	adminsIds []int,
) Chat {
	return Chat{
		id:         id,
		avatar:     avatar,
		title:      title,
		type_:      type_,
		membersIds: membersIds,
		isArchived: isArchived,
		ownerId:    ownerId,
		adminsIds:  adminsIds,
	}
}
