package bug

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// getByProfile godoc
// @Summary      Get bugs assigned to teams this profile is a part of.
// @Description  Get bugs assigned to teams this profile is a part of.
// @Tags         bug
// @Produce      json
// @Success      200  {object}  []bugOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /bug/by-profile [get]
func (handler *Handler) getByProfile(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	bugs, err := handler.store.GetActiveBugsByProfile(r.Context(), payload.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	response := []bugOutput{}
	for _, bug := range bugs {
		closedBy, _ := uuid.FromBytes(bug.Closedby.Bytes[:])
		remarks := bug.Remarks.String
		closedAt := bug.Closedat.Time

		response = append(response, bugOutput{
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

	util.WriteJson(w, http.StatusOK, response)
}
