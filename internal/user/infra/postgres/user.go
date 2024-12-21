package postgres

import (
	"context"

	"main/internal/user/domain/model"
	"main/internal/user/infra"
	"main/internal/user/infra/postgres/dto"
)

type UserRepository struct {
	db infra.DB
	//	here need logger
}

func (u UserRepository) Get(ctx context.Context) ([]model.User, error) {

	rows, err := u.db.Query(ctx, "SELECT id, username FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]dto.User, 0)

	for rows.Next() {
		var user dto.User
		if err = rows.Scan(
			&user.ID,
			&user.Username,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return dto.UsersAdapter(users), nil
}

func (u UserRepository) GetById(ctx context.Context, id model.ID) (model.User, error) {

	row := u.db.QueryRow(ctx, "SELECT id, username FROM users where id = $1", id)

	var user dto.User

	if err := row.Scan(
		&user.ID,
		&user.Username,
	); err != nil {
		return model.User{}, err
	}

	return dto.UserAdapter(user), nil
}

func (u UserRepository) GetByUsername(ctx context.Context, username string) (model.User, error) {

	row := u.db.QueryRow(ctx, "SELECT id, username FROM users where username = $1", username)

	var user dto.User

	if err := row.Scan(
		&user.ID,
		&user.Username,
	); err != nil {
		return model.User{}, err
	}

	return dto.UserAdapter(user), nil
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

func (u UserRepository) Delete(ctx context.Context, id model.ID) error {
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
		id,
	); err != nil || ct.RowsAffected() != 1 {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
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
