package profile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) createHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateProfileInput
	if err := util.ReadJsonAndValidate(w, r, &req); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	res, err := handler.Create(r.Context(), &req)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	util.WriteJson(w, http.StatusOK, res)
}

func (handler *Handler) Create(ctx context.Context, req *CreateProfileInput) (*ProfileOutput, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, util.ErrInternal
	}

	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, util.ErrInternal
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, util.ErrInternal
	}

	profile, err := handler.store.CreateProfile(ctx, db.CreateProfileParams{
		ID:       id,
		Tokenid:  tokenId,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, util.ErrEmailExists
		}
		return nil, util.ErrDatabase
	}

	accessToken, err := handler.tokenMaker.CreateToken(ctx, profile.ID, profile.Tokenid, token.EmailToken, handler.config.EMAIL_DURATION)
	if err != nil {
		return nil, util.ErrInternal
	}

	go util.SendMail(util.SendMailArgs{
		From:         handler.config.SENDER_EMAIL,
		Password:     handler.config.SENDER_PASSWORD,
		To:           profile.Email,
		Subject:      "Welcome to Buggy!",
		TemplatePath: "./verification-email.html",
		TemplateData: map[string]interface{}{
			"Name": profile.Name,
			"Link": fmt.Sprintf("%s/api/v%d/profile/verify?token=%s", handler.config.API_PREFIX, handler.config.API_VERSION, accessToken),
		},
		EmailHost:     handler.config.EMAIL_HOST,
		EmailHostPort: handler.config.EMAIL_HOST_PORT,
	})

	return &ProfileOutput{
		Id:        profile.ID,
		Name:      profile.Name,
		Email:     profile.Email,
		Verified:  profile.Verified,
		CreatedAt: profile.Createdat.Unix(),
		UpdatedAt: profile.Updatedat.Unix(),
	}, nil
}
