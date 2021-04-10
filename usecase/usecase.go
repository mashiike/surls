package usecase

import (
	"context"
	"net/url"
)

//Usecase is is a container for Boundarys
type Usecase struct {
	GetShortcutBoundary
	CreateShortcutBoundary
}

//GetShortcutBoundary is Boundary of GetShortcut
type GetShortcutBoundary interface {
	GetShortcut(context.Context, *GetShortcutInput) (*GetShortcutOutput, error)
}

//GetShortcutInput is InputData for GetShortcutBoundary
type GetShortcutInput struct {
	ShortcutID string `json:"shortcut_id"`
}

//GetShortcutInput is OutputData for GetShortcutBoundary
type GetShortcutOutput struct {
	RedirectURL *url.URL
}

//CreateShortcutBoundary is Boundary of CreateShortcut
type CreateShortcutBoundary interface {
	CreateShortcut(context.Context, *CreateShortcutInput) (*CreateShortcutOutput, error)
}

//CreateShortcutInput is InputData for CreateShortcutBoundary
type CreateShortcutInput struct {
	LongURL string `json:"long_url"`
}

//CreateShortcutInput is OutputData for CreateShortcutBoundary
type CreateShortcutOutput struct {
	ShortcutID string
	LongURL    string
}
