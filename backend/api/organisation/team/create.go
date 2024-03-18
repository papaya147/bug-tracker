package team

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// create godoc
// @Summary      Create a new team under a profile's organisation.
// @Description  Create a new team under a profile's organisation.
// @Tags         organisation
// @Accept       json
// @Produce      json
// @Param 		 input body createTeamInput true "json"
// @Success      200  {object}  teamOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /organisation/team [post]
func (handler *Handler) create(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	var requestPayload createTeamInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	org, err := handler.store.GetOrganisationByOwner(r.Context(), payload.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	teamId, err := uuid.NewV7()
	if err != nil {
		util.NewErrorAndWrite(w, util.ErrInternal)
		return
	}

	team, err := handler.store.CreateTeam(r.Context(), db.CreateTeamParams{
		ID:           teamId,
		Name:         requestPayload.Name,
		Description:  requestPayload.Description,
		Organisation: org.ID,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			util.NewErrorAndWrite(w, util.ErrEntityExists)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	go handler.store.CreateTeamMember(context.Background(), db.CreateTeamMemberParams{
		Team:    teamId,
		Profile: payload.UserId,
		Admin:   true,
	})

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
