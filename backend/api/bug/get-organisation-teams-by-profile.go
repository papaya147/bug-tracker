package bug

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// getOrganisationTeamsByProfile godoc
// @Summary      Get teams this profile is part of.
// @Description  Get teams this profile is part of.
// @Tags         bug
// @Produce      json
// @Param        organisation-id  path  string  true  "Organisation ID"
// @Success      200  {object}  []teamOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /bug/organisation/{organisation-id}/teams-by-profile [get]
func (handler *Handler) getOrganisationTeamsByProfile(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	requestPayload := id{
		Id: chi.URLParam(r, "organisation-id"),
	}
	if err := util.ValidateRequest(requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	teams, err := handler.store.GetTeamsByProfileAndOrganisation(r.Context(), db.GetTeamsByProfileAndOrganisationParams{
		Profile: payload.UserId,
		ID:      uuid.MustParse(requestPayload.Id),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	response := []teamOutput{}
	for _, team := range teams {
		response = append(response, teamOutput{
			Id:          team.Teamid,
			Name:        team.Teamname,
			Description: team.Teamdescription,
		})
	}

	util.WriteJson(w, http.StatusOK, response)
}
