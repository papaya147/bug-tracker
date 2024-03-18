package bug

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
	"github.com/papaya147/parallelize"
)

// create godoc
// @Summary      Create a new bug.
// @Description  Create a new bug, the assigned team and assignee team must be part of the same organisation and the profile must be a part of the assignee team.
// @Tags         bug
// @Accept       json
// @Produce      json
// @Param 		 input body createBugInput true "json"
// @Success      200  {object}  bugOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /bug [post]
func (handler *Handler) create(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	var requestPaylod createBugInput
	if err := util.ReadJsonAndValidate(w, r, &requestPaylod); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	if _, err := handler.store.GetTeamMember(r.Context(), db.GetTeamMemberParams{
		Team:    requestPaylod.AssigneeTeam,
		Profile: payload.UserId,
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrUnauthorised)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	var org1, org2 uuid.UUID
	group := parallelize.NewSyncGroup()
	parallelize.AddOutputtingMethodWithArgs(group, handler.getTeamOrganisation, parallelize.OutputtingMethodWithArgsParams[uuid.UUID, *uuid.UUID]{
		Context: r.Context(),
		Input:   requestPaylod.AssignedTeam,
		Output:  &org1,
	})
	parallelize.AddOutputtingMethodWithArgs(group, handler.getTeamOrganisation, parallelize.OutputtingMethodWithArgsParams[uuid.UUID, *uuid.UUID]{
		Context: r.Context(),
		Input:   requestPaylod.AssigneeTeam,
		Output:  &org2,
	})
	if err := group.Run(); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	if org1 != org2 {
		util.NewErrorAndWrite(w, util.ErrDifferentOrganisation)
		return
	}

	bugId, _ := uuid.NewV7()

	bug, err := handler.store.CreateBug(r.Context(), db.CreateBugParams{
		ID:                bugId,
		Name:              requestPaylod.Name,
		Description:       requestPaylod.Description,
		Priority:          requestPaylod.Priority,
		Assignedto:        requestPaylod.AssignedTeam,
		Assignedbyprofile: payload.UserId,
		Assignedbyteam:    requestPaylod.AssigneeTeam,
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

func (handler *Handler) getTeamOrganisation(ctx context.Context, teamId uuid.UUID, orgId *uuid.UUID) error {
	org, err := handler.store.GetTeamOrganisation(ctx, teamId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return util.ErrEntityDoesNotExist
		}
		return util.ErrDatabase
	}
	*orgId = org
	return nil
}
