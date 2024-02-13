package token

import (
	"context"
	"encoding/json"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
	"github.com/papaya147/buggy/backend/util"
)

type Maker interface {
	CreateToken(ctx context.Context, userId, tokenId uuid.UUID, typ TokenType, duration time.Duration) (string, error)
	VerifyToken(ctx context.Context, token string) (*Payload, error)
}

type PasetoMaker struct {
	parser       paseto.Parser
	symmetricKey paseto.V4SymmetricKey
}

func NewPasetoMaker(ctx context.Context, secret string) (Maker, error) {
	secretKey, err := paseto.V4SymmetricKeyFromHex(secret)
	if err != nil {
		return nil, err
	}

	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())

	maker := &PasetoMaker{
		parser:       parser,
		symmetricKey: secretKey,
	}
	return maker, nil
}

// CreateToken implements Maker.
func (maker *PasetoMaker) CreateToken(ctx context.Context, userId, tokenId uuid.UUID, typ TokenType, duration time.Duration) (string, error) {
	payload := NewPayload(userId, tokenId, typ, duration)

	token := paseto.NewToken()
	token.SetIssuedAt(payload.IssuedAt)
	token.SetNotBefore(time.Now())
	token.SetExpiration(payload.ExpiredAt)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	token.SetString("payload", string(jsonData))

	return token.V4Encrypt(maker.symmetricKey, nil), nil
}

var payload Payload

// VerifyToken implements Maker.
func (maker *PasetoMaker) VerifyToken(ctx context.Context, token string) (*Payload, error) {
	parsedToken, err := maker.parser.ParseV4Local(maker.symmetricKey, token, nil)
	if err != nil {
		return nil, util.ErrInvalidToken
	}

	p, err := parsedToken.GetString("payload")
	if err != nil {
		return nil, util.ErrInvalidToken
	}

	err = json.Unmarshal([]byte(p), &payload)
	if err != nil {
		return nil, util.ErrInternal
	}

	return &payload, nil
}
