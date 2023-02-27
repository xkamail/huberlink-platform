package auth

import (
	"time"
)

type RefreshToken struct {
	ID        int64
	UserID    int64
	Token     string
	ExpiredAt time.Time
	IssueAt   time.Time
	CreatedAt time.Time
}
