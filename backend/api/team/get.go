package team

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// get godoc
// @Summary      Get all teams for a profile.
// @Description  Get all teams for a profile.
// @Tags         team
// @Produce      json
// @Success      200  {object}  []teamOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /team [get]
func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	teams, err := h.store.GetTeams(r.Context(), payload.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	var response []teamOutput
	for _, team := range teams {
		response = append(response, teamOutput{
			Id:                      team.Teamid,
			OrganisationName:        team.Orgname,
			OrganisationDescription: team.Orgdescription,
			Name:                    team.Teamname,
			Description:             team.Teamdescription,
			Admin:                   team.Admin,
		})
	}

	util.WriteJson(w, http.StatusOK, response)
}
