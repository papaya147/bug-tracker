package bug

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// getOrganisationTeams godoc
// @Summary      Get organisations this profile is part of.
// @Description  Get organisations this profile is part of.
// @Tags         bug
// @Produce      json
// @Param        organisation-id  path  string  true  "Organisation ID"
// @Success      200  {object}  []teamOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /bug/organisation/{organisation-id}/teams [get]
func (handler *Handler) getOrganisationTeams(w http.ResponseWriter, r *http.Request) {
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

	orgId := uuid.MustParse(requestPayload.Id)

	orgs, err := handler.store.GetAssignableOrganisations(r.Context(), payload.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	validReq := false
	for _, org := range orgs {
		if org.ID == orgId {
			validReq = true
			break
		}
	}

	if !validReq {
		util.NewErrorAndWrite(w, util.ErrUnauthorised)
		return
	}

	teams, err := handler.store.GetOrganisationTeams(r.Context(), orgId)
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
			Id:          team.ID,
			Name:        team.Name,
			Description: team.Description,
			CreatedAt:   team.Createdat.Unix(),
			UpdatedAt:   team.Updatedat.Unix(),
		})
	}

	util.WriteJson(w, http.StatusOK, response)
}
