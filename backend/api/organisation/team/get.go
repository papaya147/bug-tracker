package team

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) get(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenPayloadFromContext(r.Context(), token.AccessToken)
	if err != nil {
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

	teams, err := handler.store.GetOrganisationTeams(r.Context(), org.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			util.WriteJson(w, http.StatusOK, []teamResponse{})
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	var response []teamResponse
	for _, team := range teams {
		response = append(response, teamResponse{
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
