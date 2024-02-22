package profile

import (
	"errors"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) changePassword(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenPayloadFromContext(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	var requestPayload changePasswordRequest
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
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

	if err := util.ValidatePassword(requestPayload.OldPassword, profile.Password); err != nil {
		util.NewErrorAndWrite(w, util.ErrWrongPassword)
		return
	}

	hashedNewPass, err := util.HashPassword(requestPayload.NewPassword)
	if err != nil {
		util.NewErrorAndWrite(w, util.ErrInternal)
		return
	}

	profile, err = handler.store.UpdatePassword(r.Context(), db.UpdatePasswordParams{
		Password: hashedNewPass,
		ID:       payload.UserId,
	})
	if err != nil {
		log.Println(err)
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, ProfileOutput{
		Id:        profile.ID,
		Name:      profile.Name,
		Email:     profile.Email,
		Verified:  profile.Verified,
		CreatedAt: profile.Createdat.Unix(),
		UpdatedAt: profile.Updatedat.Unix(),
	})
}
