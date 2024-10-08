package token

import (
	"time"
)

type Maker interface {
	// CreateToken create a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks if token is valid or not
	VerifyToken(token string) (*Payload, error)
}
