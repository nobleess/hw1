package model

import (
	"github.com/google/uuid"
	"main/internal/channel/channel"
	user "main/internal/user/domain/model"
	"time"
)

type ID uuid.UUID

type Message interface {
	Id() ID
	UserId() user.ID
	ChannelId() channel.ID

	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() time.Time
}
