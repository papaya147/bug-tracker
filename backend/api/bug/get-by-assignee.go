package bug

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// getByProfile godoc
// @Summary      Get bugs assigned by this team.
// @Description  Get bugs assigned by this team.
// @Tags         bug
// @Produce      json
// @Param        team-id  path  string  true  "Team ID"
// @Success      200  {object}  []bugOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /bug/by-assignee-team/{team-id} [get]
func (handler *Handler) getByAssigneeTeam(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	requestPayload := id{
		Id: chi.URLParam(r, "team-id"),
	}
	if err := util.ValidateRequest(requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	if _, err := handler.store.GetTeamMember(r.Context(), db.GetTeamMemberParams{
		Team:    uuid.MustParse(requestPayload.Id),
		Profile: payload.UserId,
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrUnauthorised)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	bugs, err := handler.store.GetBugsByAssigneeTeam(r.Context(), uuid.MustParse(requestPayload.Id))
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
