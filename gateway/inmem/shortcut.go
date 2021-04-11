package inmem

import (
	"context"

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
