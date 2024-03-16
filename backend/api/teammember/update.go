package teammember

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// create godoc
// @Summary      Update a team members permissions.
// @Description  Update a team members permissions.
// @Tags         team-member
// @Accept       json
// @Produce      json
// @Param 		 input body updateTeamMemberInput true "json"
// @Success      200  {object}  teamMemberOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /team-member [put]
func (handler *Handler) update(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	var requestPayload updateTeamMemberInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	org, err := handler.store.GetOrganisation(r.Context(), payload.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrEntityDoesNotExist)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	if org.Owner == requestPayload.ProfileId {
		util.NewErrorAndWrite(w, util.ErrUnauthorised)
		return
	}

	member, err := handler.store.GetTeamMember(r.Context(), db.GetTeamMemberParams{
		Team:    requestPayload.TeamId,
		Profile: payload.UserId,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrUnauthorised)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	if !member.Admin {
		util.NewErrorAndWrite(w, util.ErrUnauthorised)
		return
	}

	member, err = handler.store.UpdateTeamMember(r.Context(), db.UpdateTeamMemberParams{
		Admin:   requestPayload.Admin,
		Team:    requestPayload.TeamId,
		Profile: requestPayload.ProfileId,
	})
	if err != nil {
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, teamMemberOutput{
		TeamId:    member.Team,
		ProfileId: member.Profile,
		Admin:     member.Admin,
		CreatedAt: member.Createdat.Unix(),
		UpdatedAt: member.Updatedat.Unix(),
	})
}
