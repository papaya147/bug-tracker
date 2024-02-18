package profile

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) forgotPassword(w http.ResponseWriter, r *http.Request) {
	var requestPayload forgotEmailRequest
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	profile, err := handler.store.GetProfileByEmail(r.Context(), requestPayload.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			util.ErrorJson(w, util.ErrUserNotFound)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	token, err := handler.tokenMaker.CreateToken(r.Context(), profile.ID, profile.Tokenid, token.PasswordToken, handler.config.EMAIL_DURATION)
	if err != nil {
		util.ErrorJson(w, util.ErrInternal)
		return
	}

	redirectPath := fmt.Sprintf("%s%s?token=%s", handler.config.CLIENT_PREFIX, handler.config.FORGOT_PASSWORD_PATH, token)
	// TODO - send the path as a link in an email
	fmt.Println(redirectPath)
}
