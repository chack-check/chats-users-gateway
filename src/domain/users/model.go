package users

import "github.com/chack-check/chats-users-gateway/domain/files"

type UserPermissionCategory struct {
	id   int
	code string
	name string
}

func (c *UserPermissionCategory) GetId() int {
	return c.id
}

func (c *UserPermissionCategory) GetCode() string {
	return c.code
}

func (c *UserPermissionCategory) GetName() string {
	return c.name
}

type UserPermission struct {
	id       int
	code     string
	name     string
	category *UserPermissionCategory
}

func (p *UserPermission) GetId() int {
	return p.id
}

func (p *UserPermission) GetCode() string {
	return p.code
}

func (p *UserPermission) GetName() string {
	return p.name
}

func (p *UserPermission) GetCategory() *UserPermissionCategory {
	return p.category
}

type User struct {
	id             int
	username       string
	avatar         *files.SavedFile
	phone          *string
	email          *string
	firstName      string
	lastName       string
	middleName     *string
	status         *string
	emailConfirmed bool
	phoneConfirmed bool
	permissions    []UserPermission
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetAvatar() *files.SavedFile {
	return u.avatar
}

func (u *User) GetPhone() *string {
	return u.phone
}

func (u *User) GetEmail() *string {
	return u.email
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) GetMiddleName() *string {
	return u.middleName
}

func (u *User) GetStatus() *string {
	return u.status
}

func (u *User) GetEmailConfirmed() bool {
	return u.emailConfirmed
}

func (u *User) GetPhoneConfirmed() bool {
	return u.phoneConfirmed
}

func (u *User) GetPermissions() []UserPermission {
	return u.permissions
}

func NewUserPermissionCategory(id int, code, name string) UserPermissionCategory {
	return UserPermissionCategory{
		id:   id,
		code: code,
		name: name,
	}
}

func NewUserPermission(id int, code, name string, category *UserPermissionCategory) UserPermission {
	return UserPermission{
		id:       id,
		code:     code,
		name:     name,
		category: category,
	}
}

func NewUser(
	id int,
	username string,
	avatar *files.SavedFile,
	phone *string,
	email *string,
	firstName string,
	lastName string,
	middleName *string,
	status *string,
	emailConfirmed bool,
	phoneConfirmed bool,
	permissions []UserPermission,
) User {
	return User{
		id:             id,
		username:       username,
		avatar:         avatar,
		phone:          phone,
		email:          email,
		firstName:      firstName,
		lastName:       lastName,
		middleName:     middleName,
		status:         status,
		emailConfirmed: emailConfirmed,
		phoneConfirmed: phoneConfirmed,
		permissions:    permissions,
	}
}
