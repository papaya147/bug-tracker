package team

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

// update godoc
// @Summary      Update a team under a profile's organisation.
// @Description  Update a team under a profile's organisation.
// @Tags         organisation
// @Produce      json
// @Param 		 team-id path string true "Team ID"
// @Success      200  {object}  teamOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /organisation/team/{team-id} [put]
func (handler *Handler) update(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	teamId := teamId{
		Id: chi.URLParam(r, "team-id"),
	}
	if err := util.ValidateRequest(teamId); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	var requestPayload createTeamInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
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

	team, err := handler.store.UpdateTeam(r.Context(), db.UpdateTeamParams{
		Name:         requestPayload.Name,
		Description:  requestPayload.Description,
		ID:           uuid.MustParse(teamId.Id),
		Organisation: org.ID,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			util.NewErrorAndWrite(w, util.ErrEntityExists)
			return
		}
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, teamOutput{
		Id:                      team.ID,
		OrganisationName:        org.Name,
		OrganisationDescription: org.Description,
		Name:                    team.Name,
		Description:             team.Description,
		CreatedAt:               team.Createdat.Unix(),
		UpdatedAt:               team.Updatedat.Unix(),
	})
}
