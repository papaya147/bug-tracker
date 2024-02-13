package token

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type MockMaker struct {
}

// CreateToken implements Maker.
func (MockMaker) CreateToken(ctx context.Context, userId uuid.UUID, tokenId uuid.UUID, typ TokenType, duration time.Duration) (string, error) {
	return "", nil
}

// VerifyToken implements Maker.
func (MockMaker) VerifyToken(ctx context.Context, token string) (*Payload, error) {
	return &Payload{}, nil
}

func NewMockMaker() Maker {
	return MockMaker{}
}
