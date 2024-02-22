package profile

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) verify(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	payload, err := handler.tokenMaker.VerifyToken(r.Context(), token)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	profile, err := handler.store.GetProfile(r.Context(), payload.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrProfileNotFound)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	if profile.Verified {
		util.NewErrorAndWrite(w, util.ErrProfileAlreadyVerified)
		return
	}

	if profile.Tokenid != payload.TokenId {
		util.NewErrorAndWrite(w, util.ErrInvalidToken)
		return
	}

	profile, err = handler.store.VerifyProfile(r.Context(), payload.UserId)
	if err != nil {
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, profileResponse{
		Id:        profile.ID,
		Name:      profile.Name,
		Email:     profile.Email,
		Verified:  profile.Verified,
		CreatedAt: profile.Createdat.Unix(),
		UpdatedAt: profile.Updatedat.Unix(),
	})
}
