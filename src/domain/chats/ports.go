package chats

type ChatsPort interface {
	GetById(id int) *Chat
	GetByIds(ids []int) []Chat
}

type GetConcreteChatPort interface {
	Execute(handler GetConcreteChatHandler) (*ChatWithUsers, error)
}

type GetChatsByIdsPort interface {
	Execute(handler GetChatsByIdsHandler) ChatsWithUsers
}
