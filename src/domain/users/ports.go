package users

type UsersPort interface {
	GetById(id int) *User
	GetByIds(ids []int) []User
}
