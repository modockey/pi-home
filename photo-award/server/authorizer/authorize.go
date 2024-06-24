package authorizer

type User struct {
	id   int
	name string
}

func Authorize(token string) (interface{}, error) {
	return User{0, "temp"}, nil
}
