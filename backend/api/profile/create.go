package profile

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) create(w http.ResponseWriter, r *http.Request) {
	var requestPayload createProfileRequest
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	id, err := uuid.NewV7()
	if err != nil {
		util.ErrorJson(w, util.ErrInternal)
		return
	}

	tokenId, err := uuid.NewRandom()
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

	profile, err := handler.store.CreateProfile(r.Context(), db.CreateProfileParams{
		ID:       id,
		Tokenid:  tokenId,
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

	token, err := handler.tokenMaker.CreateToken(r.Context(), profile.ID, profile.Tokenid, token.EmailToken, handler.config.EMAIL_DURATION)
	if err != nil {
		util.ErrorJson(w, util.ErrInternal)
		return
	}

	go util.SendMail(util.SendMailArgs{
		From:         handler.config.SENDER_EMAIL,
		Password:     handler.config.SENDER_PASSWORD,
		To:           profile.Email,
		Subject:      "Welcome to Buggy!",
		TemplatePath: "./verification-email.html",
		TemplateData: map[string]interface{}{
			"Name": profile.Name,
			"Link": fmt.Sprintf("%s/api/v%d/profile/verify?token=%s", handler.config.API_PREFIX, handler.config.API_VERSION, token),
		},
		EmailHost:     handler.config.EMAIL_HOST,
		EmailHostPort: handler.config.EMAIL_HOST_PORT,
	})

	util.WriteJson(w, http.StatusOK, profileResponse{
		Id:        profile.ID,
		Name:      profile.Name,
		Email:     profile.Email,
		Verified:  profile.Verified,
		CreatedAt: profile.Createdat.Unix(),
		UpdatedAt: profile.Updatedat.Unix(),
	})
}
