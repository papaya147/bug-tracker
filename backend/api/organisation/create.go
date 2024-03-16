package organisation

import (
	"net/http"

	"github.com/google/uuid"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// create godoc
// @Summary      Create a new organisation.
// @Description  Create a new organisation, each profile may only have one organisation.
// @Tags         organisation
// @Accept       json
// @Produce      json
// @Param 		 input body createOrganisationInput true "json"
// @Success      200  {object}  organisationOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /organisation [post]
func (handler *Handler) create(w http.ResponseWriter, r *http.Request) {
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

	orgId, err := uuid.NewV7()
	if err != nil {
		util.NewErrorAndWrite(w, util.ErrInternal)
		return
	}

	org, err := handler.store.CreateOrganisation(r.Context(), db.CreateOrganisationParams{
		ID:          orgId,
		Name:        requestPayload.Name,
		Description: requestPayload.Description,
		Owner:       payload.UserId,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			util.NewErrorAndWrite(w, util.ErrEntityExists)
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
