package token

import (
	"time"

	"github.com/google/uuid"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	EmailToken   TokenType = "email"
	InvalidToken TokenType = "invalid"
)

func (t TokenType) String() string {
	return string(t)
}

type Payload struct {
	UserId    uuid.UUID `json:"id"`
	TokenId   uuid.UUID `json:"token_id"`
	TokenType TokenType `json:"token_type"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userId, tokenId uuid.UUID, typ TokenType, duration time.Duration) *Payload {
	return &Payload{
		UserId:    userId,
		TokenId:   tokenId,
		TokenType: typ,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}
