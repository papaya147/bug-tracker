package profile

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) login(w http.ResponseWriter, r *http.Request) {
	var requestPayload loginRequest
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	profile, err := handler.store.UpdateTokenIdByEmail(r.Context(), db.UpdateTokenIdByEmailParams{
		Tokenid: uuid.New(),
		Email:   requestPayload.Email,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			util.ErrorJson(w, util.ErrUserNotFound)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	if err := util.ValidatePassword(requestPayload.Password, profile.Password); err != nil {
		util.ErrorJson(w, util.ErrWrongPassword)
		return
	}

	token, err := handler.tokenMaker.CreateToken(r.Context(), profile.ID, profile.Tokenid, token.AccessToken, handler.config.SESSION_DURATION)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	util.WriteJson(w, http.StatusOK, profileResponse{
		Id:        profile.ID,
		Name:      profile.Name,
		Email:     profile.Email,
		Verified:  profile.Verified,
		CreatedAt: profile.Createdat.Unix(),
		UpdatedAt: profile.Updatedat.Unix(),
		Token:     token,
	})
}
