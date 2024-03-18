package bug

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// getOrganisations godoc
// @Summary      Get organisations this profile is part of.
// @Description  Get organisations this profile is part of.
// @Tags         bug
// @Produce      json
// @Success      200  {object}  []organisationOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /bug/organisations [get]
func (handler *Handler) getOrganisations(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	orgs, err := handler.store.GetAssignableOrganisations(r.Context(), payload.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	response := []organisationOutput{}
	for _, org := range orgs {
		response = append(response, organisationOutput{
			ID:          org.ID,
			Name:        org.Name,
			Description: org.Description,
			Owner:       org.Owner,
			CreatedAt:   org.Createdat.Unix(),
			UpdatedAt:   org.Updatedat.Unix(),
		})
	}

	util.WriteJson(w, http.StatusOK, response)
}
