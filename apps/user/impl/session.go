package impl

import "github.com/rs/xid"

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