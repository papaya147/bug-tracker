// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Bug struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Status      string           `json:"status"`
	Priority    string           `json:"priority"`
	Assignedto  uuid.UUID        `json:"assignedto"`
	Assignedby  uuid.UUID        `json:"assignedby"`
	Closedby    pgtype.UUID      `json:"closedby"`
	Createdat   pgtype.Timestamp `json:"createdat"`
	Updatedat   pgtype.Timestamp `json:"updatedat"`
	Closedat    pgtype.Timestamp `json:"closedat"`
}

type Organisation struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Owner       uuid.UUID `json:"owner"`
	Createdat   time.Time `json:"createdat"`
	Updatedat   time.Time `json:"updatedat"`
}

type Profile struct {
	ID        uuid.UUID `json:"id"`
	Tokenid   uuid.UUID `json:"tokenid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Verified  bool      `json:"verified"`
	Createdat time.Time `json:"createdat"`
	Updatedat time.Time `json:"updatedat"`
}

type Team struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Organisation uuid.UUID `json:"organisation"`
	Createdat    time.Time `json:"createdat"`
	Updatedat    time.Time `json:"updatedat"`
}

type Teammember struct {
	Team      uuid.UUID `json:"team"`
	Profile   uuid.UUID `json:"profile"`
	Createdat time.Time `json:"createdat"`
	Updatedat time.Time `json:"updatedat"`
}
