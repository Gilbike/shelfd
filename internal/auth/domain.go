package auth

import "time"

type Session struct {
	ID        string    `json:"id"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
	UserAgent string    `json:"user_agent"`
	IpAddress string    `json:"ip_address"`
}

func (session *Session) IsValid() bool {
	return time.Now().Before(session.ExpiresAt)
}
