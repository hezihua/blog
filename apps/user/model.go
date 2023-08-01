package user

func NewSession(username, sess string) *Session {
	return &Session{
		Username: username,
		Id:       sess,
	}
}

type Session struct {
	Username string `json:"username"`
	Id       string `json:"session"`
}