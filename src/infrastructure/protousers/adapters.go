package protousers

import (
	"context"

	"github.com/chack-check/chats-users-gateway/domain/files"
	"github.com/chack-check/chats-users-gateway/domain/users"
	"github.com/chack-check/chats-users-gateway/infrastructure/protousers/usersprotobuf"
)

type UsersAdapter struct {
	Client usersprotobuf.UsersClient
}

func (adapter *UsersAdapter) GetById(id int) *users.User {
	user, err := adapter.Client.GetUserById(context.Background(), &usersprotobuf.GetUserByIdRequest{Id: int32(id)})
	if err != nil {
		return nil
	}

	var avatar *files.SavedFile
	if user.Avatar != nil {
		file := files.NewSavedFile(
			user.Avatar.OriginalUrl,
			user.Avatar.OriginalFilename,
			user.Avatar.ConvertedUrl,
			user.Avatar.ConvertedFilename,
		)
		avatar = &file
	}

	domainUser := users.NewUser(
		int(user.Id),
		user.Username,
		avatar,
		user.Phone,
		user.Email,
		user.FirstName,
		user.LastName,
		user.MiddleName,
		user.Status,
		user.EmailConfirmed,
		user.PhoneConfirmed,
		[]users.UserPermission{},
	)
	return &domainUser
}

func (adapter *UsersAdapter) GetByIds(ids []int) []users.User {
	var ids32 []int32
	for _, id := range ids {
		ids32 = append(ids32, int32(id))
	}

	fetchedUsers, err := adapter.Client.GetUsersByIds(context.Background(), &usersprotobuf.GetUsersByIdsRequest{Ids: ids32})
	if err != nil {
		return []users.User{}
	}

	var domainUsers []users.User
	for _, user := range fetchedUsers.Users {
		var avatar *files.SavedFile
		if user.Avatar != nil {
			file := files.NewSavedFile(
				user.Avatar.OriginalUrl,
				user.Avatar.OriginalFilename,
				user.Avatar.ConvertedUrl,
				user.Avatar.ConvertedFilename,
			)
			avatar = &file
		}

		domainUser := users.NewUser(
			int(user.Id),
			user.Username,
			avatar,
			user.Phone,
			user.Email,
			user.FirstName,
			user.LastName,
			user.MiddleName,
			user.Status,
			user.EmailConfirmed,
			user.PhoneConfirmed,
			[]users.UserPermission{},
		)
		domainUsers = append(domainUsers, domainUser)
	}

	return domainUsers
}
