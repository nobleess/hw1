package postgres

import (
	"context"
	"main/internal/message/infra"
	"main/internal/message/infra/postgres/dto"
	"main/internal/user/domain/model"
)

type UserRepository struct {
	db infra.DB
	//	somelse
}

func (u UserRepository) GetUsers(ctx context.Context) ([]model.User, error) {

	rows, err := u.db.Query(ctx, "SELECT id, username FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]dto.User, 0)

	for rows.Next() {
		var u dto.User
		if err = rows.Scan(
			&u.ID,
			&u.Username,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return dto.UserAdapter(users), nil
}

func (u UserRepository) Create(ctx context.Context, user model.User) error {

	tx, err := u.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if ct, err := tx.Exec(
		ctx,
		"insert into model values ()",
		user.ID(),
		user.Login(),
	); err != nil || ct.RowsAffected() != 1 {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil
	}
	return nil
}

func (u UserRepository) Delete(ctx context.Context, user model.User) error {
	tx, err := u.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if ct, err := tx.Exec(
		ctx,
		"delete from model where id = $1",
		user.ID(),
	); err != nil || ct.RowsAffected() != 1 {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil
	}
	return nil
}

func (u UserRepository) Update(ctx context.Context, user model.User) error {
	tx, err := u.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if ct, err := tx.Exec(
		ctx,
		"update model set login=$2 where id = $1",
		user.ID(),
		user.Login(),
	); err != nil || ct.RowsAffected() != 1 {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil
	}
	return nil
}
