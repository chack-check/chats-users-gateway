package messages

import (
	"time"

	"github.com/chack-check/chats-users-gateway/domain/files"
	"github.com/chack-check/chats-users-gateway/domain/users"
)

type PaginatedMessages struct {
	Offset   int
	Limit    int
	Total    int
	Messages []Message
}

type PaginatedMessagesWithUsers struct {
	Messages []Message
	Users    []users.User
	Limit    int
	Offset   int
	Total    int
}

type MessageReaction struct {
	id      int
	userId  int
	content string
}

func (r *MessageReaction) GetId() int {
	return r.id
}

func (r *MessageReaction) GetUserId() int {
	return r.userId
}

func (r *MessageReaction) GetContent() string {
	return r.content
}

type Message struct {
	id          int
	senderId    int
	chatId      int
	type_       string
	content     string
	voice       *files.SavedFile
	circle      *files.SavedFile
	attachments []files.SavedFile
	replyToId   *int
	mentioned   []int
	readedBy    []int
	reactions   []MessageReaction
	deletedFor  []int
	createdAt   *time.Time
}

func (m *Message) GetId() int {
	return m.id
}

func (m *Message) GetSenderId() int {
	return m.senderId
}

func (m *Message) GetChatId() int {
	return m.chatId
}

func (m *Message) GetType() string {
	return m.type_
}

func (m *Message) GetContent() string {
	return m.content
}

func (m *Message) GetVoice() *files.SavedFile {
	return m.voice
}

func (m *Message) GetCircle() *files.SavedFile {
	return m.circle
}

func (m *Message) GetAttachments() []files.SavedFile {
	return m.attachments
}

func (m *Message) GetReplyToId() *int {
	return m.replyToId
}

func (m *Message) GetMentioned() []int {
	return m.mentioned
}

func (m *Message) GetReadedBy() []int {
	return m.readedBy
}

func (m *Message) GetReaction() []MessageReaction {
	return m.reactions
}

func (m *Message) GetDeletedFor() []int {
	return m.deletedFor
}

func (m *Message) GetCreatedAt() *time.Time {
	return m.createdAt
}

func NewMessage(
	id int,
	senderId int,
	chatId int,
	type_ string,
	content string,
	voice *files.SavedFile,
	circle *files.SavedFile,
	attachments []files.SavedFile,
	replyToId *int,
	mentioned []int,
	readedBy []int,
	reactions []MessageReaction,
	deletedFor []int,
	createdAt *time.Time,
) Message {
	return Message{
		id:          id,
		senderId:    senderId,
		chatId:      chatId,
		type_:       type_,
		content:     content,
		voice:       voice,
		circle:      circle,
		attachments: attachments,
		replyToId:   replyToId,
		mentioned:   mentioned,
		readedBy:    readedBy,
		reactions:   reactions,
		deletedFor:  deletedFor,
		createdAt:   createdAt,
	}
}

func NewMessageReaction(id int, userId int, content string) MessageReaction {
	return MessageReaction{
		id:      id,
		userId:  userId,
		content: content,
	}
}
