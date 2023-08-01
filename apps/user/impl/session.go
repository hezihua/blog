package impl

import (
	"context"
	"fmt"
	"hezihua/apps/user"

	"github.com/rs/xid"
)

func (i *impl) createSession(username string) string {
	i.lock.Lock()
	defer i.lock.Unlock()

	sess := xid.New().String()
	i.Sessions[username] = sess
	return sess
	// return nil, nil
}

func (i *impl) deleteSession(username string) string {
	i.lock.Lock()
	defer i.lock.Unlock()
	
	sess := xid.New().String()

	delete(i.Sessions, username)
	return sess
	// return nil, nil
}

func (i *impl) CheckIsLogin(ctx context.Context, req *user.CheckIsLoginRequest) error {
	sessId, ok := i.Sessions[req.Username]; 
	if !ok {
		return fmt.Errorf("user %s not found", req.Username)
	}
	if req.SessionId != sessId {
		return fmt.Errorf("session %s not found", req.Username)
	}
	return nil

}