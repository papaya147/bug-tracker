package transfer

import "github.com/google/uuid"

type transferRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type transferResponse struct {
	Id           uuid.UUID `json:"id"`
	Organisation uuid.UUID `json:"organisation"`
	FromProfile  uuid.UUID `json:"from_profile"`
	ToProfile    uuid.UUID `json:"to_profile"`
	Completed    bool      `json:"completed"`
	CreatedAt    int64     `json:"created_at"`
}

type transferDetailedResponse struct {
	Id                      uuid.UUID `json:"id"`
	OrganisationName        uuid.UUID `json:"organisation_name"`
	OrganisationDescription string    `json:"organisation_description"`
	FromProfile             uuid.UUID `json:"from_profile"`
	ToProfile               uuid.UUID `json:"to_profile"`
	Completed               bool      `json:"completed"`
	CreatedAt               int64     `json:"created_at"`
}

type transferResponseList struct {
	IncomingTransfers []transferDetailedResponse `json:"incoming_transfers"`
	OutgoingTransfer  *transferDetailedResponse  `json:"outgoing_transfer"`
}
