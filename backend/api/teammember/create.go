package teammember

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func (handler *Handler) create(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenPayloadFromContext(r.Context(), token.AccessToken)
	if err != nil {
		util.ErrorJson(w, err)
		return
	}

	var requestPayload createTeamMemberRequest
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	member, err := handler.store.GetTeamMember(r.Context(), db.GetTeamMemberParams{
		Team:    requestPayload.TeamId,
		Profile: payload.UserId,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			util.ErrorJson(w, util.ErrUnauthorised)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	if !member.Admin {
		util.ErrorJson(w, util.ErrUnauthorised)
		return
	}

	profile, err := handler.store.GetProfileByEmail(r.Context(), requestPayload.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			util.ErrorJson(w, util.ErrProfileNotFound)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	newMember, err := handler.store.CreateTeamMember(r.Context(), db.CreateTeamMemberParams{
		Team:    requestPayload.TeamId,
		Profile: profile.ID,
		Admin:   requestPayload.Admin,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			util.ErrorJson(w, util.ErrTeamMemberAlreadyExists)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, teamMemberResponse{
		TeamId:    newMember.Team,
		ProfileId: newMember.Profile,
		Admin:     newMember.Admin,
		CreatedAt: newMember.Createdat.Unix(),
		UpdatedAt: newMember.Updatedat.Unix(),
	})
}
