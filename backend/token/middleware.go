package token

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/util"
)

type TokenPayloadType string

var TokenPayloadKey TokenPayloadType = "token-payload"

func Middleware(tokenMaker Maker, store db.Store) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				util.NewErrorAndWrite(w, util.ErrInvalidToken)
				return
			}

			if len(auth) <= len("Bearer ") {
				util.NewErrorAndWrite(w, util.ErrInvalidToken)
				return
			}

			bearer := "Bearer "
			auth = auth[len(bearer):]

			payload, err := tokenMaker.VerifyToken(r.Context(), auth)
			if err != nil {
				util.NewErrorAndWrite(w, err)
				return
			}

			profile, err := store.GetProfile(r.Context(), payload.UserId)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					util.NewErrorAndWrite(w, util.ErrProfileNotFound)
					return
				}
				util.NewErrorAndWrite(w, util.ErrDatabase)
				return
			}

			if profile.Tokenid != payload.TokenId {
				util.NewErrorAndWrite(w, util.ErrInvalidToken)
				return
			}

			ctx := context.WithValue(r.Context(), TokenPayloadKey, *payload)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func GetTokenPayloadFromContext(ctx context.Context, tokenType TokenType) (Payload, error) {
	tokenPayload := ctx.Value(TokenPayloadKey).(Payload)
	if tokenPayload.TokenType != tokenType {
		return Payload{}, util.ErrInvalidToken
	}
	return tokenPayload, nil
}
