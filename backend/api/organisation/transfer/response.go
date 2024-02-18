package transfer

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) response(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenPayloadFromContext(r.Context(), token.AccessToken)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	requestPayload := transferResponseStatus{
		Id:     chi.URLParam(r, "organisation-transfer-id"),
		Status: r.URL.Query().Get("status"),
	}

	if err := util.ValidateRequest(requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	stat, _ := strconv.ParseBool(requestPayload.Status)

	if stat {
		if _, err = handler.store.CompleteOrganisationTransferTx(r.Context(), db.CompleteOrganisationTransferTxParams{
			TransferId: uuid.MustParse(requestPayload.Id),
			ToProfile:  payload.UserId,
		}); err != nil {
			util.ErrorJson(w, err)
			return
		}
	} else {
		if _, err = handler.store.DeleteOrganisationTransfer(r.Context(), uuid.MustParse(requestPayload.Id)); err != nil {
			if err == pgx.ErrNoRows {
				util.ErrorJson(w, util.ErrEntityDoesNotExist)
				return
			}
			util.ErrorJson(w, util.ErrDatabase)
			return
		}
	}

	util.WriteJson(w, http.StatusOK, nil)
}
