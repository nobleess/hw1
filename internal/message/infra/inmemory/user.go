package inmemory

import (
	"errors"
	"main/internal/message/domain/model/user"
)

type UserStorage struct {
	storage map[user.Login]*user.User
}

func NewUserStorage(storage map[user.Login]*user.User) *UserStorage {
	return &UserStorage{
		storage: storage,
	}
}

func (s *UserStorage) GetUsers() []*user.User {
	users := make([]*user.User, 0, len(s.storage))
	i := 0
	for _, u := range s.storage {
		users[i] = u
		i++
	}
	return users
}

func (s *UserStorage) FindUserByLogin(login user.Login) (*user.User, error) {
	if u, ok := s.storage[login]; ok {
		return nil, nil
	} else {
		return u, errors.New("user not found")
	}
}

func (s *UserStorage) FindUserById(id user.ID) *user.User {
	for _, u := range s.storage {
		if u.ID() == id {
			return u
		}
	}
	return nil
}
