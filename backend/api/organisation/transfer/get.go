package transfer

import (
	"net/http"

	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) get(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenPayloadFromContext(r.Context(), token.AccessToken)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	transfers, err := handler.store.GetOrganisationTransfersTx(r.Context(), payload.UserId)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	util.WriteJson(w, http.StatusOK, transfers)
}
