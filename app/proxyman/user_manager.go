package proxyman

import (
	"sync"
	"xray-core/common/protocol"
)

// UserManager manages the list of users.
type UserManager struct {
	sync.Mutex
	users []*protocol.User
}

// RemoveExpiredUsers removes users whose expiration date has passed.
func (m *UserManager) RemoveExpiredUsers() {
	m.Lock()
	defer m.Unlock()

	var activeUsers []*protocol.User
	for _, user := range m.users {
		if !user.IsExpired() {
			activeUsers = append(activeUsers, user)
		}
	}
	m.users = activeUsers
}

func NewUserManager() *UserManager {
	manager := &UserManager{
		users: make([]*protocol.User, 0),
	}

	go func() {
		for {
			manager.RemoveExpiredUsers()
			time.Sleep(1 * time.Hour) // Jalankan setiap 1 jam
		}
	}()

	return manager
}
