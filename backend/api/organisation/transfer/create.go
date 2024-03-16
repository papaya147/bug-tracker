package transfer

import (
	"net/http"

	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// create godoc
// @Summary      Create a new organisational transfer.
// @Description  Create a new organisational transfer.
// @Tags         organisation
// @Accept       json
// @Produce      json
// @Param 		 input body transferInput true "json"
// @Success      200  {object}  transferOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /organisation/transfer [post]
func (handler *Handler) create(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	var requestPayload transferInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	transfer, err := handler.store.CreateOrganisationTransferTx(r.Context(), db.CreateOrganisationTransferTxParams{
		FromProfile: payload.UserId,
		ToEmail:     requestPayload.Email,
	})
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	util.WriteJson(w, http.StatusOK, transferOutput{
		Id:           transfer.ID,
		Organisation: transfer.Organisation,
		FromProfile:  transfer.Fromprofile,
		ToProfile:    transfer.Toprofile,
		Completed:    transfer.Completed,
		CreatedAt:    transfer.Createdat.Unix(),
	})
}
