package organisation

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) update(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenPayloadFromContext(r.Context(), token.AccessToken)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	var requestPayload createOrganisationRequest
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	org, err := handler.store.UpdateOrganisation(r.Context(), db.UpdateOrganisationParams{
		Owner:       payload.UserId,
		Name:        requestPayload.Name,
		Description: requestPayload.Description,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			util.ErrorJson(w, util.ErrEntityDoesNotExist)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, organisationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		Owner:       org.Owner,
		CreatedAt:   org.Createdat.Unix(),
		UpdatedAt:   org.Updatedat.Unix(),
	})
}
