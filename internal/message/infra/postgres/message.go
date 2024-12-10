package postgres

import (
	"context"
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/user"
	"main/internal/message/infra"
	"main/internal/message/infra/postgres/dto"
)

type MessageReposytory struct {
	db infra.DB
}

func NewMessageReposytory(db infra.DB) *MessageReposytory {
	return &MessageReposytory{
		db: db,
	}
}

func (m MessageReposytory) Create(ctx context.Context, msg message.Message) error {

	tx, err := m.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	if ct, err := tx.Exec(
		ctx,
		// todo
		"insert into message values ()",
	); err != nil || ct.RowsAffected() != 1 {
		return err
	}

	err = tx.Commit(ctx)
	return err
}

func (m MessageReposytory) Get(ctx context.Context, id user.ID) ([]dto.Message, error) {

}

func (m MessageReposytory) FindByUserId(ctx context.Context, id user.ID) ([]dto.Message, error) {

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

func (m MessageReposytory) FindByChannelId(ctx context.Context, id user.ChannelId) ([]dto.Message, error) {

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
