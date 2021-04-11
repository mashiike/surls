package inmem

import (
	"context"
	"errors"

	"github.com/mashiike/surls/entity"
	"github.com/mashiike/surls/usecase"
)

type ShortcutRepository struct {
	store map[entity.ShortcutID]*entity.Shortcut
}

func NewShortcutRepository() usecase.ShortcutRepository {
	return &ShortcutRepository{
		store: make(map[entity.ShortcutID]*entity.Shortcut),
	}
}

func (repo *ShortcutRepository) Save(_ context.Context, shortcut *entity.Shortcut) error {
	repo.store[shortcut.ID()] = shortcut
	return nil
}

func (repo *ShortcutRepository) Find(_ context.Context, id entity.ShortcutID) (*entity.Shortcut, error) {
	shortcut, ok := repo.store[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return shortcut, nil
}
