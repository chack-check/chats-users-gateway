package protochats

import (
	"context"
	"log"
	"time"

	"github.com/chack-check/chats-users-gateway/domain/chats"
	"github.com/chack-check/chats-users-gateway/domain/files"
	"github.com/chack-check/chats-users-gateway/domain/messages"
	protobuf "github.com/chack-check/chats-users-gateway/infrastructure/protochats/chatsprotobuf"
	"github.com/chack-check/chats-users-gateway/utils"
)

func protoSavedFileToSavedFile(file *protobuf.SavedFile) *files.SavedFile {
	var convertedUrl *string
	var convertedFilename *string
	if file.GetConvertedUrl() != "" {
		convertedUrlValue := file.GetConvertedUrl()
		convertedFilenameValue := file.GetConvertedFilename()
		convertedUrl = &convertedUrlValue
		convertedFilename = &convertedFilenameValue
	} else {
		convertedUrl = nil
		convertedFilename = nil
	}

	savedFile := files.NewSavedFile(
		file.GetOriginalUrl(),
		file.GetOriginalFilename(),
		convertedUrl,
		convertedFilename,
	)
	return &savedFile
}

func protoChatToChat(chat *protobuf.ChatResponse) *chats.Chat {
	var chatAvatar *files.SavedFile
	if avatar := chat.GetAvatar(); avatar != nil {
		chatAvatar = protoSavedFileToSavedFile(avatar)
	} else {
		chatAvatar = nil
	}

	newChat := chats.NewChat(
		int(chat.GetId()),
		chatAvatar,
		chat.GetTitle(),
		chat.GetType(),
		utils.ConvertArrayItems[int32, int](chat.GetMembersIds()),
		chat.GetIsArchived(),
		int(chat.GetOwnerId()),
		utils.ConvertArrayItems[int32, int](chat.GetAdminsIds()),
	)
	return &newChat
}

func protoReactionToMessageReaction(reaction *protobuf.MessageReaction) *messages.MessageReaction {
	newReaction := messages.NewMessageReaction(
		int(reaction.GetId()),
		int(reaction.GetUserId()),
		reaction.GetContent(),
	)
	return &newReaction
}

func protoMessageToMessage(message *protobuf.MessageResponse) *messages.Message {
	voiceFile := protoSavedFileToSavedFile(message.GetVoice())
	circleFile := protoSavedFileToSavedFile(message.GetCircle())
	var attachments []files.SavedFile
	for _, attachment := range message.GetAttachments() {
		attachments = append(attachments, *protoSavedFileToSavedFile(attachment))
	}

	var reactions []messages.MessageReaction
	for _, reaction := range message.GetReactions() {
		reactions = append(reactions, *protoReactionToMessageReaction(reaction))
	}

	var replyToId *int
	if message.GetReplyToId() != 0 {
		replyToIdInt := int(message.GetReplyToId())
		replyToId = &replyToIdInt
	} else {
		replyToId = nil
	}

	var createdAtDt *time.Time
	if message.CreatedAt != nil {
		timeObject, err := time.Parse(time.RFC3339, *message.CreatedAt)
		if err == nil {
			createdAtDt = &timeObject
		}
	}

	newMessage := messages.NewMessage(
		int(message.GetId()),
		int(message.GetSenderId()),
		int(message.GetChatId()),
		message.GetType(),
		message.GetContent(),
		voiceFile,
		circleFile,
		attachments,
		replyToId,
		utils.ConvertArrayItems[int32, int](message.GetMentioned()),
		utils.ConvertArrayItems[int32, int](message.GetReadedBy()),
		reactions,
		utils.ConvertArrayItems[int32, int](message.GetDeletedFor()),
		createdAtDt,
	)

	return &newMessage
}

type ChatsAdapter struct {
	chats.ChatsPort

	Client protobuf.ChatsClient
	Token  string
}

func (adapter *ChatsAdapter) GetById(id int) *chats.Chat {
	chatResponse, err := adapter.Client.GetChatById(context.Background(), &protobuf.GetChatByIdRequest{Id: int32(id), Token: adapter.Token})
	if err != nil {
		return nil
	}

	return protoChatToChat(chatResponse)
}

func (adapter *ChatsAdapter) GetByIds(ids []int) []chats.Chat {
	convertedIds := utils.ConvertArrayItems[int, int32](ids)
	chatsResponses, err := adapter.Client.GetChatsByIds(context.Background(), &protobuf.GetChatsByIdsRequest{Ids: convertedIds, Token: adapter.Token})
	if err != nil {
		return []chats.Chat{}
	}

	var chats []chats.Chat
	for _, chatResponse := range chatsResponses.GetChats() {
		chat := protoChatToChat(chatResponse)
		if chat == nil {
			continue
		}

		chats = append(chats, *chat)
	}

	return chats
}

type MessagesAdapter struct {
	messages.MessagesPort

	Client protobuf.ChatsClient
	Token  string
}

func (adapter *MessagesAdapter) GetById(id int) *messages.Message {
	messageResponse, err := adapter.Client.GetMessageById(context.Background(), &protobuf.GetMessageByIdRequest{Id: int32(id), Token: adapter.Token})
	if err != nil {
		return nil
	}

	return protoMessageToMessage(messageResponse)
}

func (adapter *MessagesAdapter) GetByChatId(chatId int, offset int, limit int) messages.PaginatedMessages {
	log.Printf("Sending token: %s", adapter.Token)
	offset32 := int32(offset)
	limit32 := int32(limit)
	messagesResponse, err := adapter.Client.GetMessagesByChatId(
		context.Background(),
		&protobuf.GetMessagesByChatIdRequest{
			ChatId: int32(chatId),
			Offset: &offset32,
			Limit:  &limit32,
			Token:  adapter.Token,
		},
	)
	log.Printf("messages response: %+v. Token: %s. Error: %v", messagesResponse, adapter.Token, err)
	if err != nil {
		return messages.PaginatedMessages{
			Limit:    0,
			Offset:   1,
			Total:    0,
			Messages: []messages.Message{},
		}
	}

	var messagesModels []messages.Message
	for _, message := range messagesResponse.Data {
		messagesModels = append(messagesModels, *protoMessageToMessage(message))
	}

	return messages.PaginatedMessages{
		Limit:    int(messagesResponse.Limit),
		Offset:   int(messagesResponse.Offset),
		Total:    int(messagesResponse.Total),
		Messages: messagesModels,
	}
}
