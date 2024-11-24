package service

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/ports"
)

type PlayerServiceOpt func(s *PlayerService)

func WithTransaction(tx ports.DatabaseTxPort) PlayerServiceOpt {
	return func(s *PlayerService) {
		s.tx = tx
	}
}

type PlayerService struct {
	db ports.DatabasePort
	tx ports.DatabaseTxPort
}

func NewPlayerService(database ports.DatabasePort, opts ...PlayerServiceOpt) (*PlayerService, error) {
	service := &PlayerService{
		db: database,
	}

	for _, opt := range opts {
		opt(service)
	}

	return service, nil
}

func (s *PlayerService) NewPlayerTxService(ctx context.Context) (ports.PlayerTxServicePort, error) {
	txAdapter, err := s.db.NewTransactionAdapter(ctx)
	if err != nil {
		return nil, err
	}

	txService, err := NewPlayerService(txAdapter, WithTransaction(txAdapter))
	if err != nil {
		return nil, err
	}

	return txService, nil
}

func (s *PlayerService) Commit(ctx context.Context) error {
	return s.tx.Commit(ctx)
}

func (s *PlayerService) Rollback(ctx context.Context) error {
	return s.tx.Rollback(ctx)
}
