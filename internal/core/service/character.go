package service

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/core/entity"
	"github.com/nxdir-s/idlerpg/internal/ports"
)

type CharacterServiceOpt func(s *CharacterService)

func WithCharacterTx(tx ports.DatabaseTx) CharacterServiceOpt {
	return func(s *CharacterService) {
		s.tx = tx
	}
}

type CharacterService struct {
	db ports.Database
	tx ports.DatabaseTx
}

func NewCharacterService(ctx context.Context, database ports.Database, opts ...CharacterServiceOpt) (*CharacterService, error) {
	service := &CharacterService{
		db: database,
	}

	for _, opt := range opts {
		opt(service)
	}

	return service, nil
}

func (s *CharacterService) NewCharacterTxService(ctx context.Context) (ports.CharacterTxService, error) {
	txAdapter, err := s.db.NewTransactionAdapter(ctx)
	if err != nil {
		return nil, err
	}

	txService, err := NewCharacterService(ctx, txAdapter, WithCharacterTx(txAdapter))
	if err != nil {
		return nil, err
	}

	return txService, nil
}

func (s *CharacterService) Commit(ctx context.Context) error {
	return s.tx.Commit(ctx)
}

func (s *CharacterService) Rollback(ctx context.Context) error {
	return s.tx.Rollback(ctx)
}

func (s *CharacterService) GetCharacter(ctx context.Context, userId int) (*entity.Character, error) {
	return s.db.GetCharacter(ctx, userId)
}
