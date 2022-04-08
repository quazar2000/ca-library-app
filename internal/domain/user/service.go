package user

import "context"

type Service interface {
	GetUserByUUID(ctx context.Context, uuid string)
	GetAllUsers(ctx context.Context, limit, offset int)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) GetUserByUUID(ctx context.Context, uuid string) *User {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAllUsers(ctx context.Context, limit, offset int) []*User {
	return s.storage.GetAll(limit, offset)
}
