package messages

type MessagesPort interface {
	GetById(id int) *Message
	GetByChatId(chatId int, offset int, limit int) PaginatedMessages
}

type GetChatMessagesPort interface {
	Execute(handler GetChatMessagesHandler) PaginatedMessagesWithUsers
}
