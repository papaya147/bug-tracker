package profile

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/util"
)

// verify godoc
// @Summary      Verify a profile.
// @Description  Verify a profile. This link will only come from emails.
// @Tags         profile
// @Produce      json
// @Param 		 token query string true "string"
// @Success      200  {object}  profileOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /profile/verify [get]
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

	http.Redirect(w, r, fmt.Sprintf("%s/login", handler.config.CLIENT_PREFIX), http.StatusFound)
}
