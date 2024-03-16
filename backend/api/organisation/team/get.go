package team

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// create godoc
// @Summary      Get all teams under a profile's organisation.
// @Description  Get all teams under a profile's organisation.
// @Tags         organisation
// @Produce      json
// @Success      200  {object}  []teamOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /organisation/team [get]
func (handler *Handler) get(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	org, err := handler.store.GetOrganisation(r.Context(), payload.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	teams, err := handler.store.GetOrganisationTeams(r.Context(), org.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.WriteJson(w, http.StatusOK, []teamOutput{})
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	var response []teamOutput
	for _, team := range teams {
		response = append(response, teamOutput{
			Id:                      team.ID,
			OrganisationName:        org.Name,
			OrganisationDescription: org.Description,
			Name:                    team.Name,
			Description:             team.Description,
			CreatedAt:               team.Createdat.Unix(),
			UpdatedAt:               team.Updatedat.Unix(),
		})
	}

	util.WriteJson(w, http.StatusOK, response)
}
