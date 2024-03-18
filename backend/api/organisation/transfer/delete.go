package transfer

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// delete godoc
// @Summary      Delete an organisational transfer.
// @Description  Delete an organisational transfer.
// @Tags         organisation
// @Accept       json
// @Produce      json
// @Param 		 organisation-transfer-id path string true "Organisation Transfer ID"
// @Success      200  {object}  transferOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /organisation/transfer/{organisation-transfer-id} [delete]
func (handler *Handler) delete(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	requestPayload := transferId{
		Id: chi.URLParam(r, "organisation-transfer-id"),
	}

	if err := util.ValidateRequest(requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	_, err = handler.store.GetOrganisationByOwner(r.Context(), payload.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	transfer, err := handler.store.DeleteOrganisationTransfer(r.Context(), uuid.MustParse(requestPayload.Id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
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
