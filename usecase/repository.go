package usecase

import (
	"context"

	"github.com/mashiike/surls/entity"
)

type ShortcutRepository interface {
	Save(context.Context, *entity.Shortcut) error
}
