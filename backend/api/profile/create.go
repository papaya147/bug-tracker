package profile

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	db "github.com/papaya147/bug-tracker/backend/db/sqlc"
	"github.com/papaya147/bug-tracker/backend/util"
)

func (handler *Handler) create(w http.ResponseWriter, r *http.Request) {
	var requestPayload createProfileInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		util.ErrorJson(w, util.ErrInternal)
		return
	}

	hashedPassword, err := util.HashPassword(requestPayload.Password)
	if err != nil {
		log.Println(err)
		util.ErrorJson(w, util.ErrInternal)
		return
	}

	profile, err := handler.store.Queries.CreateProfile(r.Context(), db.CreateProfileParams{
		ID:       id,
		Name:     requestPayload.Name,
		Email:    requestPayload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			util.ErrorJson(w, util.ErrEmailExists)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	// TODO - add email verification

	util.WriteJson(w, http.StatusOK, profileOutput{
		Id:        profile.ID,
		Name:      profile.Name,
		Email:     profile.Email,
		Verified:  profile.Verified,
		CreatedAt: profile.Createdat.Unix(),
		UpdatedAt: profile.Updatedat.Unix(),
	})
}
