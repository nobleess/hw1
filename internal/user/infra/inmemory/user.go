package inmemory

// trash
//import (
//	"errors"
//	"main/internal/user/domain/model"
//)
//
//type UserStorage struct {
//	storage map[model.Login]*model.User
//}
//
//func NewUserStorage(storage map[model.Login]*model.User) *UserStorage {
//	return &UserStorage{
//		storage: storage,
//	}
//}
//
//func (s *UserStorage) GetUsers() []*model.User {
//	users := make([]*model.User, 0, len(s.storage))
//	i := 0
//	for _, u := range s.storage {
//		users[i] = u
//		i++
//	}
//	return users
//}
//
//func (s *UserStorage) FindUserByLogin(login model.Login) (*model.User, error) {
//	if u, ok := s.storage[login]; ok {
//		return nil, nil
//	} else {
//		return u, errors.New("model not found")
//	}
//}
//
//func (s *UserStorage) FindUserById(id model.ID) *model.User {
//	for _, u := range s.storage {
//		if u.ID() == id {
//			return u
//		}
//	}
//	return nil
//}
