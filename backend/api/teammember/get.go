package teammember

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

func (handler *Handler) get(w http.ResponseWriter, r *http.Request) {
	payload, err := token.GetTokenPayloadFromContext(r.Context(), token.AccessToken)
	if err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	requestPayload := teamId{
		Id: chi.URLParam(r, "team-id"),
	}

	if err := util.ValidateRequest(requestPayload); err != nil {
		util.NewErrorAndWrite(w, err)
		return
	}

	if _, err = handler.store.GetTeamMember(r.Context(), db.GetTeamMemberParams{
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

	profiles, err := handler.store.GetAllTeamMembers(r.Context(), uuid.MustParse(requestPayload.Id))
	if err != nil {
		util.NewErrorAndWrite(w, util.ErrDatabase)
		return
	}

	var response []profileResponse
	for _, member := range profiles {
		response = append(response, profileResponse{
			Id:        member.ID,
			Admin:     member.Admin,
			Name:      member.Name,
			Email:     member.Email,
			Verified:  member.Verified,
			CreatedAt: member.Createdat.Unix(),
			UpdatedAt: member.Updatedat.Unix(),
		})
	}

	util.WriteJson(w, http.StatusOK, response)
}
