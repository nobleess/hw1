package postgres

import (
	"context"

	"main/internal/message/domain/model/user"
	"main/internal/message/infra"
	"main/internal/message/infra/dto"
	"main/internal/message/infra/postgres/dto"
)

type UserRepository struct {
	db infra.DB
	//	somelse
}

func (u UserRepository) GetUsers(ctx context.Context) ([]dto.User, error) {

	rows, err := u.db.Query(ctx, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]dto.User, 0)

	for rows.Next() {
		var u dto.user
		if err = rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Data,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepository) Create(ctx context.Context, user user.User) error {

	tx, err := u.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if ct, err := tx.Exec(
		ctx,
		"insert into user values ()",
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

func (u UserRepository) Delete(ctx context.Context, user user.User) error {
	tx, err := u.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if ct, err := tx.Exec(
		ctx,
		"delete from user where id = $1",
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

func (u UserRepository) Update(ctx context.Context, user user.User) error {
	tx, err := u.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if ct, err := tx.Exec(
		ctx,
		"update user set login=$2 where id = $1",
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
