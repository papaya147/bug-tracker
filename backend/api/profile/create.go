package profile

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// create godoc
// @Summary      Create a new profile, emails must be unique.
// @Description  Create a new profile by giving a profile name, email and password. This API will generate a verification email which will contain a redirect link.
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param 		 input body createProfileInput true "json"
// @Success      200  {object}  profileOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /profile [post]
func (handler *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req createProfileInput
	if err := util.ReadJsonAndValidate(w, r, &req); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	id, err := uuid.NewV7()
	if err != nil {
		util.NewErrorAndWrite(w, util.ErrInternal)
	}

	tokenId, err := uuid.NewRandom()
	if err != nil {
		util.NewErrorAndWrite(w, util.ErrInternal)
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		util.NewErrorAndWrite(w, util.ErrInternal)
		return
	}

	profile, err := handler.store.CreateProfile(r.Context(), db.CreateProfileParams{
		ID:       id,
		Tokenid:  tokenId,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			util.NewErrorAndWrite(w, util.ErrEmailExists)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	accessToken, err := handler.tokenMaker.CreateToken(r.Context(), profile.ID, profile.Tokenid, token.EmailToken, handler.config.EMAIL_DURATION)
	if err != nil {
		util.NewErrorAndWrite(w, util.ErrInternal)
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
			"Link": fmt.Sprintf("%s/api/v%s/profile/verify?token=%s", handler.config.API_PREFIX, handler.config.API_VERSION, accessToken),
		},
		EmailHost:     handler.config.EMAIL_HOST,
		EmailHostPort: handler.config.EMAIL_HOST_PORT,
	})

	util.WriteJson(w, http.StatusOK, profileOutput{
		Id:        profile.ID,
		Name:      profile.Name,
		Email:     profile.Email,
		Verified:  profile.Verified,
		CreatedAt: profile.Createdat.Unix(),
		UpdatedAt: profile.Updatedat.Unix(),
	})
}
