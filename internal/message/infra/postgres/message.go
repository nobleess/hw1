package postgres

import (
	"context"
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/user"
	"main/internal/message/infra"
	"main/internal/message/infra/postgres/dto"
	user2 "main/internal/user/domain/model"
)

type MessageRepository struct {
	db infra.DB
}

func NewMessageReposytory(db infra.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (m MessageRepository) Create(ctx context.Context, msg message.Message) error {

	tx, err := m.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	if ct, err := tx.Exec(
		ctx,
		// todo
		"insert into message values ($1, $2, $3, $4, $5) returning id )",
		msg.Id(),
	); err != nil || ct.RowsAffected() != 1 {
		return err
	}

	err = tx.Commit(ctx)
	return err
}

func (m MessageRepository) FindByUserId(ctx context.Context, id user2.ID) ([]dto.Message, error) {

	rows, err := m.db.Query(
		ctx,
		// todo change *
		"select * from message where user_id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	msgs := make([]dto.Message, 0)

	for rows.Next() {
		var msg dto.Message
		if err = rows.Scan(
			&msg.UserId,
			&msg.ChannelId,
			&msg.CreateAt,
			&msg.UpdateAt,
			&msg.DeleteAt,
			&msg.Text,
		); err != nil {
			return nil, err
		}
		if msg.DeleteAt == msg.CreateAt {
			msgs = append(msgs) // not deleted : deleted condition !=
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return msgs, err
}

func (m MessageRepository) FindByChannelId(ctx context.Context, id user.ChannelId) ([]dto.Message, error) {

	rows, err := m.db.Query(
		ctx,
		// todo change *
		"select * from message where channel_id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	msgs := make([]dto.Message, 0)

	for rows.Next() {
		var msg dto.Message
		if err = rows.Scan(
			&msg.UserId,
			&msg.ChannelId,
			&msg.CreateAt,
			&msg.UpdateAt,
			&msg.DeleteAt,
			&msg.Text,
		); err != nil {
			return nil, err
		}
		if msg.DeleteAt == msg.CreateAt {
			msgs = append(msgs)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return msgs, err
}
