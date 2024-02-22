package organisation

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) update(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenPayloadFromContext(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	var requestPayload createOrganisationRequest
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

	util.WriteJson(w, http.StatusOK, organisationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		Owner:       org.Owner,
		CreatedAt:   org.Createdat.Unix(),
		UpdatedAt:   org.Updatedat.Unix(),
	})
}
