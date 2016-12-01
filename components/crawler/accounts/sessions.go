package accounts

import (
	"sync"
)

var sessions = struct {
	sync.RWMutex
	m []string
}{}

func initSessions(length uint) {
	sessions.m = make([]string, length)
}

func GetSessionByIndex(index uint) string {
	sessions.RLock()
	session := sessions.m[index]
	sessions.RUnlock()
	return session
}

func UpdateSessionByIndex(index uint, session string) {
	sessions.Lock()
	sessions.m[index] = session
	sessions.Unlock()
}
