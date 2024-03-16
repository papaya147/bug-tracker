package profile

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// get godoc
// @Summary      Get a profile from a token.
// @Description  Get a profile from a token.
// @Tags         profile
// @Produce      json
// @Success      200  {object}  profileOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /profile [get]
func (handler *Handler) get(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
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

	util.WriteJson(w, http.StatusOK, profileOutput{
		Id:        profile.ID,
		Name:      profile.Name,
		Email:     profile.Email,
		Verified:  profile.Verified,
		CreatedAt: profile.Createdat.Unix(),
		UpdatedAt: profile.Updatedat.Unix(),
	})
}
