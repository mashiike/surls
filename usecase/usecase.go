package usecase

import (
	"context"

	"github.com/mashiike/surls/entity"
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
	RedirectURL string
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
	ShortcutID entity.ShortcutID
	LongURL    string
}
