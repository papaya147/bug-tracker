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
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sess, _ := util.Store.Get(r, "buggy-session")
			token, ok := sess.Values["token"].(string)
			if !ok {
				util.NewErrorAndWrite(w, util.ErrInvalidCookie)
				return
			}

			payload, err := tokenMaker.VerifyToken(r.Context(), token)
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
			next.ServeHTTP(w, r)
		})
	}
}

func GetTokenDetail(ctx context.Context, tokenType TokenType) (Payload, error) {
	tokenPayload := ctx.Value(TokenPayloadKey).(Payload)
	if tokenPayload.TokenType != tokenType {
		return Payload{}, util.ErrInvalidToken
	}
	return tokenPayload, nil
}
