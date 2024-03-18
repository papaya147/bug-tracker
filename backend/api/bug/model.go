package bug

import (
	"time"

	"github.com/google/uuid"
	db "github.com/papaya147/buggy/backend/db/sqlc"
)

type createBugInput struct {
	Name         string         `json:"name" validate:"required,max=32" example:"Improper Input Validation"`
	Description  string         `json:"description" validate:"required,max=255" example:"Input validation is not working"`
	Priority     db.Bugpriority `json:"priority" validate:"required,oneof=URGENT HIGH LOW" example:"URGENT"`
	AssignedTeam uuid.UUID      `json:"assigned_team" validate:"required" example:"00000000-0000-0000-0000-000000000000"`
	AssigneeTeam uuid.UUID      `json:"assignee_team" validate:"required" example:"00000000-0000-0000-0000-000000000000"`
}

type bugOutput struct {
	Id                uuid.UUID      `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Name              string         `json:"name" example:"Improper Input Validation"`
	Description       string         `json:"description" example:"Input validation is not working"`
	Status            db.Bugstatus   `json:"status" example:"PENDING"`
	Priority          db.Bugpriority `json:"priority" example:"URGENT"`
	Assignedto        uuid.UUID      `json:"assignedto" example:"00000000-0000-0000-0000-000000000000"`
	Assignedbyprofile uuid.UUID      `json:"assignedbyprofile" example:"00000000-0000-0000-0000-000000000000"`
	Assignedbyteam    uuid.UUID      `json:"assignedbyteam" example:"00000000-0000-0000-0000-000000000000"`
	Completed         bool           `json:"completed" example:"false"`
	Createdat         time.Time      `json:"createdat" example:"1710579130"`
	Updatedat         time.Time      `json:"updatedat" example:"1710579130"`
	Closedby          *uuid.UUID     `json:"closedby" example:"00000000-0000-0000-0000-000000000000"`
	Remarks           *string        `json:"remarks" example:"None"`
	Closedat          *time.Time     `json:"closedat" example:"1710579130"`
}

type bugId struct {
	Id string `json:"id" validate:"required,uuid" example:"00000000-0000-0000-0000-000000000000"`
}

type updateBugInput struct {
	Id          uuid.UUID      `json:"id" validate:"required" example:"00000000-0000-0000-0000-000000000000"`
	Name        string         `json:"name" validate:"max=32" example:"Improper Input Validation"`
	Description string         `json:"description" validate:"max=255" example:"Input validation is not working"`
	Status      db.Bugstatus   `json:"status" validate:"oneof=PENDING PROCESSING" example:"PENDING"`
	Priority    db.Bugpriority `json:"priority" validate:"oneof=URGENT HIGH LOW" example:"URGENT"`
}
