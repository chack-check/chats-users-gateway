package graph

import (
	"time"

	"github.com/chack-check/chats-users-gateway/domain/files"
	"github.com/chack-check/chats-users-gateway/domain/messages"
	"github.com/chack-check/chats-users-gateway/domain/users"
	"github.com/chack-check/chats-users-gateway/infrastructure/api/graph/model"
)

func DomainSavedFileToResponse(file files.SavedFile) model.SavedFile {
	return model.SavedFile{
		OriginalURL:       file.GetOriginalUrl(),
		OriginalFilename:  file.GetOriginalFilename(),
		ConvertedURL:      file.GetConvertedUrl(),
		ConvertedFilename: file.GetConvertedFilename(),
	}
}

func DomainPermissionCategoryToResponse(category users.UserPermissionCategory) model.UserPermissionCategory {
	return model.UserPermissionCategory{
		Name: category.GetName(),
		Code: category.GetCode(),
	}
}

func DomainPermissionToResponse(permission users.UserPermission) model.UserPermission {
	var category *model.UserPermissionCategory
	if c := permission.GetCategory(); c != nil {
		c_response := DomainPermissionCategoryToResponse(*c)
		category = &c_response
	}

	return model.UserPermission{
		Name:     permission.GetName(),
		Code:     permission.GetCode(),
		Category: category,
	}
}

func DomainUserToResponse(user users.User) model.User {
	var avatar *model.SavedFile
	if file := user.GetAvatar(); file != nil {
		fileResponse := DomainSavedFileToResponse(*file)
		avatar = &fileResponse
	}

	var permissions []*model.UserPermission
	for _, perm := range user.GetPermissions() {
		permResponse := DomainPermissionToResponse(perm)
		permissions = append(permissions, &permResponse)
	}

	return model.User{
		ID:             user.GetId(),
		Username:       user.GetUsername(),
		Phone:          user.GetPhone(),
		Email:          user.GetEmail(),
		FirstName:      user.GetFirstName(),
		LastName:       user.GetLastName(),
		MiddleName:     user.GetMiddleName(),
		Status:         user.GetStatus(),
		EmailConfirmed: user.GetEmailConfirmed(),
		PhoneConfirmed: user.GetPhoneConfirmed(),
		Avatar:         avatar,
		Permissions:    permissions,
	}
}

func DomainReactionToResponse(reaction messages.MessageReaction) model.Reaction {
	return model.Reaction{
		Content: reaction.GetContent(),
		UserID:  reaction.GetUserId(),
	}
}

func DomainMessageToResponse(message messages.Message) model.Message {
	senderId := message.GetSenderId()
	content := message.GetContent()
	var voice *model.SavedFile
	if file := message.GetVoice(); file != nil {
		fileResponse := DomainSavedFileToResponse(*file)
		voice = &fileResponse
	}

	var circle *model.SavedFile
	if file := message.GetCircle(); file != nil {
		fileResponse := DomainSavedFileToResponse(*file)
		circle = &fileResponse
	}

	var attachments []*model.SavedFile
	for _, attachment := range message.GetAttachments() {
		attachmentResponse := DomainSavedFileToResponse(attachment)
		attachments = append(attachments, &attachmentResponse)
	}

	var reactions []*model.Reaction
	for _, reaction := range message.GetReaction() {
		reactionResponse := DomainReactionToResponse(reaction)
		reactions = append(reactions, &reactionResponse)
	}

	var createdAtIsodt string
	if dt := message.GetCreatedAt(); dt != nil {
		createdAtIsodt = dt.Format(time.RFC3339)
	}
	return model.Message{
		ID:          message.GetId(),
		Type:        model.MessageType(message.GetType()),
		SenderID:    &senderId,
		ChatID:      message.GetChatId(),
		Content:     &content,
		Voice:       voice,
		Circle:      circle,
		ReplyToID:   message.GetReplyToId(),
		ReadedBy:    message.GetReadedBy(),
		Reactions:   reactions,
		Attachments: attachments,
		Mentioned:   message.GetMentioned(),
		Datetime:    createdAtIsodt,
	}
}

func DomainPaginatedMessagesWithUsersToResponse(paginatedMessages messages.PaginatedMessagesWithUsers) model.PaginatedMessagesWithUsers {
	var users []*model.User
	var messages []*model.Message
	for _, user := range paginatedMessages.Users {
		userResponse := DomainUserToResponse(user)
		users = append(users, &userResponse)
	}
	for _, message := range paginatedMessages.Messages {
		messageResponse := DomainMessageToResponse(message)
		messages = append(messages, &messageResponse)
	}
	return model.PaginatedMessagesWithUsers{
		Offset:   paginatedMessages.Offset,
		Limit:    paginatedMessages.Limit,
		Total:    paginatedMessages.Total,
		Messages: messages,
		Users:    users,
	}
}
