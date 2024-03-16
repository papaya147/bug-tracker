package transfer

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// response godoc
// @Summary      Respond to an organisation transfer.
// @Description  Respond to an organisation transfer.
// @Tags         organisation
// @Produce      json
// @Param 		 organisation-transfer-id path string true "Organisation Transfer ID"
// @Param 		 status query string true "Status"
// @Success      200  {object}  nil
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /organisation/transfer/reponse/{organisation-transfer-id} [get]
func (handler *Handler) response(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	requestPayload := transferResponseInput{
		Id:     chi.URLParam(r, "organisation-transfer-id"),
		Status: r.URL.Query().Get("status"),
	}

	if err := util.ValidateRequest(requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	stat, _ := strconv.ParseBool(requestPayload.Status)

	if stat {
		if _, err = handler.store.CompleteOrganisationTransferTx(r.Context(), db.CompleteOrganisationTransferTxParams{
			TransferId: uuid.MustParse(requestPayload.Id),
			ToProfile:  payload.UserId,
		}); err != nil {
			util.NewErrorAndWrite(w, err)
			return
		}
	} else {
		if _, err = handler.store.DeleteOrganisationTransfer(r.Context(), uuid.MustParse(requestPayload.Id)); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
				return
			}
			util.NewErrorAndWrite(w, util.ErrDatabase)
			return
		}
	}

	util.WriteJson(w, http.StatusOK, nil)
}
