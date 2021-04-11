package usecase

import (
	"context"
	"errors"
)

type CreateShortcutInteractor struct{}

func NewCreateShortcutInteractor() CreateShortcutBoundary {
	return &CreateShortcutInteractor{}
}

func (i *CreateShortcutInteractor) CreateShortcut(ctx context.Context, input *CreateShortcutInput) (*CreateShortcutOutput, error) {
	return nil, errors.New("not implemented yet")
}
