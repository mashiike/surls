package usecase

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mashiike/surls/entity"
)

type CreateShortcutInteractor struct {
	factory *entity.Factory
}

func NewCreateShortcutInteractor(f *entity.Factory) CreateShortcutBoundary {
	return &CreateShortcutInteractor{
		factory: f,
	}
}

func (i *CreateShortcutInteractor) CreateShortcut(_ context.Context, input *CreateShortcutInput) (*CreateShortcutOutput, error) {
	lurl, err := url.Parse(input.LongURL)
	if err != nil {
		return nil, fmt.Errorf("can not parse as URL %w", err)
	}
	s, err := i.factory.NewShortcut(lurl)
	if err != nil {
		return nil, fmt.Errorf("can not create shortcut %w", err)
	}
	output := &CreateShortcutOutput{
		ShortcutID: s.ID(),
		LongURL:    s.LongURL().String(),
	}
	return output, nil
}
