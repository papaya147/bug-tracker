package transfer

import (
	"net/http"

	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// get godoc
// @Summary      Get all organisation transfers.
// @Description  Get all incoming and outgoing organisation transfers.
// @Tags         organisation
// @Produce      json
// @Success      200  {object}  db.GetOrganisationTransfersTxResponse
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /organisation/transfer [get]
func (handler *Handler) get(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	transfers, err := handler.store.GetOrganisationTransfersTx(r.Context(), payload.UserId)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	util.WriteJson(w, http.StatusOK, transfers)
}
