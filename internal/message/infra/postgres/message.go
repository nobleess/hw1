package postgres

import (
	"context"
	. "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"main/internal/channel/channel"

	"main/internal/message/domain/model"
	"main/internal/message/infra"
	"main/internal/message/infra/postgres/dto"
	user "main/internal/user/domain/model"

	"main/pkg/sq"
	"main/tool/datetime"
)

type MessageRepository struct {
	db infra.DB
}

func NewMessageRepository(db infra.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (m MessageRepository) Create(ctx context.Context, msg dto.Message) error {

	tx, err := m.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	if tg, err := tx.Exec(
		ctx,
		"insert into messages (user_id, channel_id, text, media_type, url) values ($1, $2, $3, $4, $5)",
		msg.UserId,
		msg.ChannelId,
		msg.Text,
		msg.MediaType,
		msg.URL,
	); err != nil || tg.RowsAffected() != 1 {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m MessageRepository) FindByUserId(ctx context.Context, id user.ID) ([]model.Message, error) {

	sql, _, err := sq.Psql.Select("id, user_id, channel_id, create_at, update_at, delete_at, text").
		From("message").
		Where(Eq{"user_id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}
	rows, err := m.db.Query(ctx, sql)

	if err != nil {
		return nil, err
	}

	msgs, err := pgx.CollectRows[dto.Message](rows, func(row pgx.CollectableRow) (dto.Message, error) {
		var msg dto.Message
		err = rows.Scan(
			&msg.UserId,
			&msg.ChannelId,
			&msg.Text,
			&msg.CreateAt,
			&msg.UpdateAt,
			&msg.DeleteAt,
		)
		return msg, err
	})

	if err != nil {
		return nil, err
	}

	return dto.MessagesAdapter(msgs)
}

func (m MessageRepository) FindByChannelId(ctx context.Context, id channel.ID) ([]model.Message, error) {

	sql, _, err := sq.Psql.Select("id, user_id, channel_id, create_at, update_at, delete_at, text").
		From("message").
		Where(Eq{"channel_id": id}).
		ToSql()

	rows, err := m.db.Query(
		ctx, sql,
	)
	if err != nil {
		return nil, err
	}

	msgs, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (dto.Message, error) {
		var msg dto.Message
		err := rows.Scan(
			&msg.UserId,
			&msg.ChannelId,
			&msg.Text,
			&msg.CreateAt,
			&msg.UpdateAt,
			&msg.DeleteAt,
		)
		return msg, err
	})
	if err != nil {
		return nil, err
	}
	return dto.MessagesAdapter(msgs)
}

func (m MessageRepository) Update(ctx context.Context, msg dto.Message) error {
	tx, err := m.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	sql, args, err := sq.Psql.Update("message").
		SetMap(
			map[string]interface{}{
				"media_type": msg.MediaType,
				"url":        msg.URL,
				"text":       msg.Text,
				"updateAt":   datetime.Now(),
			}).
		Where(Eq{"id": msg.ID}).
		ToSql()

	if err != nil {
		return err
	}

	if tg, err := m.db.Exec(ctx, sql, args...); err != nil || tg.Update() != true {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
