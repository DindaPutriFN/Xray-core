package proxyman

import (
	"time"
	"log"
)

// User defines a generic user with an expiration date.
type User struct {
	Email   string
	Level   uint32
	Expired string // Format tanggal "YYYY-MM-DD"
}

// IsExpired checks if the user is expired based on the Expired field.
func (u *User) IsExpired() bool {
	if u.Expired == "" {
		return false
	}

	expiredDate, err := time.Parse("2006-01-02", u.Expired)
	if err != nil {
		log.Printf("Invalid expired date format for user %s: %v", u.Email, err)
		return false
	}

	return time.Now().After(expiredDate)
}
