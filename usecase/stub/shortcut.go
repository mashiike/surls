package stub

import (
	"context"
	"errors"
	"net/url"

	"github.com/mashiike/surls/usecase"
)

//Sutb Dummy Data Const
const (
	DummyShortcutID = "xxxxxx"
	FaildLongURL    = "https://example.org/failed"
)

//GetShortcutInteractor is stub implementation for usecase.GetShortcutInteractor
type GetShortcutInteractor struct{}

func NewGetShortcutInteractor() usecase.GetShortcutBoundary {
	return &GetShortcutInteractor{}
}

func (i *GetShortcutInteractor) GetShortcut(_ context.Context, _ *usecase.GetShortcutInput) (*usecase.GetShortcutOutput, error) {
	u, err := url.Parse("https://example.org")
	if err != nil {
		return nil, err
	}
	return &usecase.GetShortcutOutput{
		RedirectURL: u,
	}, nil
}

//CreateShortcutInteractor is stub implementation for usecase.CreateShortcutBoundary
type CreateShortcutInteractor struct{}

func NewCreateShortcutInteractor() usecase.CreateShortcutBoundary {
	return &CreateShortcutInteractor{}
}

func (i *CreateShortcutInteractor) CreateShortcut(_ context.Context, input *usecase.CreateShortcutInput) (*usecase.CreateShortcutOutput, error) {
	if input.LongURL == FaildLongURL {
		return nil, errors.New("stub failed url")
	}
	u, err := url.Parse(input.LongURL)
	if err != nil {
		return nil, err
	}
	return &usecase.CreateShortcutOutput{
		ShortcutID: "xxxxxx",
		LongURL:    u.String(),
	}, nil
}
