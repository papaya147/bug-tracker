package team

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) create(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenPayloadFromContext(r.Context(), token.AccessToken)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	var requestPayload createTeamRequest
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	org, err := handler.store.GetOrganisation(r.Context(), payload.UserId)
	if err != nil {
		if err == pgx.ErrNoRows {
			util.ErrorJson(w, util.ErrEntityDoesNotExist)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	teamId, err := uuid.NewV7()
	if err != nil {
		util.ErrorJson(w, util.ErrInternal)
		return
	}

	team, err := handler.store.CreateTeam(r.Context(), db.CreateTeamParams{
		ID:           teamId,
		Name:         requestPayload.Name,
		Description:  requestPayload.Description,
		Organisation: org.ID,
	})
	if err != nil {
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, teamResponse{
		Id:                      team.ID,
		OrganisationName:        org.Name,
		OrganisationDescription: org.Description,
		Name:                    team.Name,
		Description:             team.Description,
		CreatedAt:               team.Createdat.Unix(),
		UpdatedAt:               team.Updatedat.Unix(),
	})
}
