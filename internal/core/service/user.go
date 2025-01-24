package service

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/ports"
)

type UserServiceOpt func(s *UserService)

func WithUserTx(tx ports.DatabaseTxPort) UserServiceOpt {
	return func(s *UserService) {
		s.tx = tx
	}
}

type UserService struct {
	db ports.DatabasePort
	tx ports.DatabaseTxPort
}

func NewUserService(ctx context.Context, database ports.DatabasePort, opts ...UserServiceOpt) (*UserService, error) {
	service := &UserService{
		db: database,
	}

	for _, opt := range opts {
		opt(service)
	}

	return service, nil
}

func (s *UserService) NewUserTxService(ctx context.Context) (ports.UserTxServicePort, error) {
	txAdapter, err := s.db.NewTransactionAdapter(ctx)
	if err != nil {
		return nil, err
	}

	txService, err := NewUserService(ctx, txAdapter, WithUserTx(txAdapter))
	if err != nil {
		return nil, err
	}

	return txService, nil
}

func (s *UserService) Commit(ctx context.Context) error {
	return s.tx.Commit(ctx)
}

func (s *UserService) Rollback(ctx context.Context) error {
	return s.tx.Rollback(ctx)
}

func (s *UserService) GetUserID(ctx context.Context, email string) (int, error) {
	return s.db.GetUserID(ctx, email)
}

func (s *UserService) EmailExists(ctx context.Context, email string) (bool, error) {
	return s.db.EmailExists(ctx, email)
}

func (s *UserService) CreateUser(ctx context.Context, email string) (int, error) {
	return s.db.CreateUser(ctx, email)
}
