package inmem

import (
	"context"
	"errors"

	"github.com/mashiike/surls/entity"
	"github.com/mashiike/surls/usecase"
)

type ShortcutRepository struct {
}

func NewShortcutRepository() usecase.ShortcutRepository {
	return &ShortcutRepository{}
}

func (repo *ShortcutRepository) Save(_ context.Context, shortcut *entity.Shortcut) error {
	//Noop
	return nil
}

func (repo *ShortcutRepository) Find(_ context.Context, id entity.ShortcutID) (*entity.Shortcut, error) {
	//Noop
	return nil, errors.New("not implement yet")
}
