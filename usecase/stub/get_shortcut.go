package stub

import (
	"context"
	"net/url"

	"github.com/mashiike/surls/usecase"
)

//GetShortcutInteractor is stub implementation for usecase.GetShortcutInteractor
type GetShortcutInteractor struct{}

func NewGetShortcutInteractor() usecase.GetShortcutInteractor {
	return &GetShortcutInteractor{}
}

func (i *GetShortcutInteractor) GetShortcut(_ context.Context, input *usecase.GetShortcutInput) (*usecase.GetShortcutOutput, error) {
	u, err := url.Parse("https://example.org")
	if err != nil {
		return nil, err
	}
	return &usecase.GetShortcutOutput{
		RedirectURL: u,
	}, nil
}
