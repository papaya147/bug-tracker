package bug

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
// @Summary      Delete a bug.
// @Description  Delete a bug, the profile must be a part of the assigned or assignee teams.
// @Tags         bug
// @Produce      json
// @Param 		 bug-id  path  string  true  "Bug ID"
// @Success      200  {object}  bugOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /bug/{bug-id} [delete]
func (handler *Handler) delete(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	requestPayload := id{
		Id: chi.URLParam(r, "bug-id"),
	}
	if err := util.ValidateRequest(requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	bugId := uuid.MustParse(requestPayload.Id)

	bug, err := handler.store.GetBug(r.Context(), bugId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	if err := handler.checkBugPermissions(r.Context(), payload.UserId, bug.Assignedbyteam, uuid.Nil); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	bug, err = handler.store.DeleteBug(r.Context(), bugId)
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
