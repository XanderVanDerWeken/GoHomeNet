package users

type UserRegisteredEvent struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
}

func (e UserRegisteredEvent) Name() string {
	return "UserRegistered"
}
