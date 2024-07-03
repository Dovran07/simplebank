package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrTokenExpired = errors.New("Token is expired")
	ErrInvalidToken = errors.New("Invalid token")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (paload *Payload) Valid() error {
	if time.Now().After(paload.ExpiredAt) {
		return ErrTokenExpired
	}
	return nil
}
