package channel

import (
	"github.com/google/uuid"
	"main/internal/user/domain/model"
)

type ID uuid.UUID

type Channel struct {
	ID   ID
	Name string
}

type ChannelMembers struct {
	ID        ID
	ChannelID ID
	UserID    model.ID
}
