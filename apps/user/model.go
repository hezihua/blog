package user

func NewSession(username, sess string) *Session {
	return &Session{
		Username: username,
		Id:       sess,
	}
}

type Session struct {
	Username string
	Id       string
}