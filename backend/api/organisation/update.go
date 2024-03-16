package organisation

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// update godoc
// @Summary      Update an organisations name or description.
// @Description  Update an organisations name or description.
// @Tags         organisation
// @Accept       json
// @Produce      json
// @Param 		 input body createOrganisationInput true "json"
// @Success      200  {object}  organisationOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /organisation [put]
func (handler *Handler) update(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	var requestPayload createOrganisationInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	org, err := handler.store.UpdateOrganisation(r.Context(), db.UpdateOrganisationParams{
		Owner:       payload.UserId,
		Name:        requestPayload.Name,
		Description: requestPayload.Description,
	})
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
