package organisation

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// get godoc
// @Summary      Get an organisation.
// @Description  Get an organisation, this API will return an error if the organisation does not exist.
// @Tags         organisation
// @Produce      json
// @Success      200  {object}  organisationOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /organisation [get]
func (handler *Handler) get(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
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

	util.WriteJson(w, http.StatusOK, organisationOutput{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		Owner:       org.Owner,
		CreatedAt:   org.Createdat.Unix(),
		UpdatedAt:   org.Updatedat.Unix(),
	})
}
