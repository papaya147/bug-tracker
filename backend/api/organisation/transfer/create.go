package transfer

import (
	"net/http"

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

	var requestPayload transferRequest
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	transfer, err := handler.store.CreateOrganisationTransferTx(r.Context(), db.CreateOrganisationTransferTxParams{
		FromProfile: payload.UserId,
		ToEmail:     requestPayload.Email,
	})
	if err != nil {
		util.ErrorJson(w, err)
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
