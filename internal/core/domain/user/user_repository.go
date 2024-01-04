package user

type UserRepository interface {
	GetUserByUserName(userName string) (*User, error)
}
