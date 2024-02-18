package transfer

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) delete(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenPayloadFromContext(r.Context(), token.AccessToken)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	requestPayload := transferId{
		Id: chi.URLParam(r, "organisation-transfer-id"),
	}

	if err := util.ValidateRequest(requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	_, err = handler.store.GetOrganisation(r.Context(), payload.UserId)
	if err != nil {
		if err == pgx.ErrNoRows {
			util.ErrorJson(w, util.ErrEntityDoesNotExist)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	transfer, err := handler.store.DeleteOrganisationTransfer(r.Context(), uuid.MustParse(requestPayload.Id))
	if err != nil {
		if err == pgx.ErrNoRows {
			util.ErrorJson(w, util.ErrEntityDoesNotExist)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, transferResponse{
		Id:           transfer.ID,
		Organisation: transfer.Organisation,
		FromProfile:  transfer.Fromprofile,
		ToProfile:    transfer.Toprofile,
		Completed:    transfer.Completed,
		CreatedAt:    transfer.Createdat.Unix(),
	})
}
