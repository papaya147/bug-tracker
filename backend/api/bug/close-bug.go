package bug

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// close godoc
// @Summary      Close a bug.
// @Description  Close a bug, the profile closing the bug must be a part of the assigned team.
// @Tags         bug
// @Accept       json
// @Produce      json
// @Param 		 input body closeBugInput true "json"
// @Success      200  {object}  bugOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /bug/close [post]
func (handler *Handler) close(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	var requestPayload closeBugInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	bug, err := handler.store.GetBug(r.Context(), requestPayload.Id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	if err := handler.checkBugPermissions(r.Context(), payload.UserId, uuid.Nil, bug.Assignedto); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	bug, err = handler.store.CloseBug(r.Context(), db.CloseBugParams{
		Closedby: pgtype.UUID{Bytes: payload.UserId, Valid: true},
		Remarks:  pgtype.Text{String: requestPayload.Remarks, Valid: true},
		ID:       requestPayload.Id,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	closedBy, _ := uuid.FromBytes(bug.Closedby.Bytes[:])
	remarks := bug.Remarks.String
	closedAt := bug.Closedat.Time

	util.WriteJson(w, http.StatusOK, bugOutput{
		Id:                bug.ID,
		Name:              bug.Name,
		Description:       bug.Description,
		Status:            bug.Status,
		Priority:          bug.Priority,
		Assignedto:        bug.Assignedto,
		Assignedbyprofile: bug.Assignedbyprofile,
		Assignedbyteam:    bug.Assignedbyteam,
		Completed:         bug.Completed,
		Createdat:         bug.Createdat,
		Updatedat:         bug.Updatedat,
		Closedby:          &closedBy,
		Remarks:           &remarks,
		Closedat:          &closedAt,
	})
}
