package usecase

import (
	"context"
	"net/url"
)

//GetShortcutInteractor is Interactor of GetShortcut
type GetShortcutInteractor interface {
	GetShortcut(context.Context, *GetShortcutInput) (*GetShortcutOutput, error)
}

//GetShortcutInput is InputData for GetShortcutInteractor
type GetShortcutInput struct {
	ShortcutID string
}

//GetShortcutInput is OutputData for GetShortcutInteractor
type GetShortcutOutput struct {
	RedirectURL *url.URL
}
