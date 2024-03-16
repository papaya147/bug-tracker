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
// @Summary      Create a new team member.
// @Description  Create a new team member. The profile and team has to exist already to add a member. The profile making this request must be an admin for the team already.
// @Tags         team-member
// @Accept       json
// @Produce      json
// @Param 		 input body createTeamMemberInput true "json"
// @Success      200  {object}  teamMemberOutput
// @Failure      400  {object}  util.ErrorModel
// @Failure      404  {object}  util.ErrorModel
// @Failure      500  {object}  util.ErrorModel
// @Router       /team-member [post]
func (handler *Handler) create(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenDetail(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	var requestPayload createTeamMemberInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
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

	profile, err := handler.store.GetProfileByEmail(r.Context(), requestPayload.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.NewErrorAndWrite(w, util.ErrProfileNotFound)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	newMember, err := handler.store.CreateTeamMember(r.Context(), db.CreateTeamMemberParams{
		Team:    requestPayload.TeamId,
		Profile: profile.ID,
		Admin:   requestPayload.Admin,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			util.NewErrorAndWrite(w, util.ErrTeamMemberAlreadyExists)
			return
		}
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, teamMemberOutput{
		TeamId:    newMember.Team,
		ProfileId: newMember.Profile,
		Admin:     newMember.Admin,
		CreatedAt: newMember.Createdat.Unix(),
		UpdatedAt: newMember.Updatedat.Unix(),
	})
}
