package usecase

import (
	"context"
	"net/url"
)

//Usecase is is a container for Interactors
type Usecase struct {
	GetShortcutInteractor
	CreateShortcutInteractor
}

//GetShortcutInteractor is Interactor of GetShortcut
type GetShortcutInteractor interface {
	GetShortcut(context.Context, *GetShortcutInput) (*GetShortcutOutput, error)
}

//GetShortcutInput is InputData for GetShortcutInteractor
type GetShortcutInput struct {
	ShortcutID string `json:"shortcut_id"`
}

//GetShortcutInput is OutputData for GetShortcutInteractor
type GetShortcutOutput struct {
	RedirectURL *url.URL
}

//CreateShortcutInteractor is Interactor of CreateShortcut
type CreateShortcutInteractor interface {
	CreateShortcut(context.Context, *CreateShortcutInput) (*CreateShortcutOutput, error)
}

//CreateShortcutInput is InputData for CreateShortcutInteractor
type CreateShortcutInput struct {
	LongURL string `json:"long_url"`
}

//CreateShortcutInput is OutputData for CreateShortcutInteractor
type CreateShortcutOutput struct {
	ShortcutID string
	LongURL    string
}
