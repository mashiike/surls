package usecase

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mashiike/surls/entity"
)

type GetShortcutInteractor struct {
	factory *entity.Factory
	repo    ShortcutRepository
}

func NewGetShortcutInteractor(f *entity.Factory, repo ShortcutRepository) GetShortcutBoundary {
	return &GetShortcutInteractor{
		factory: f,
		repo:    repo,
	}
}

func (i *GetShortcutInteractor) GetShortcut(ctx context.Context, input *GetShortcutInput) (*GetShortcutOutput, error) {
	id, err := i.factory.NewShortcutID(input.ShortcutID)
	if err != nil {
		return nil, fmt.Errorf("ShortcutID is invalid %w", err)
	}
	shortcut, err := i.repo.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Find Shortcut failed %w", err)
	}
	output := &GetShortcutOutput{
		RedirectURL: shortcut.LongURL().String(),
	}
	return output, nil
}

type CreateShortcutInteractor struct {
	factory *entity.Factory
	repo    ShortcutRepository
}

func NewCreateShortcutInteractor(f *entity.Factory, repo ShortcutRepository) CreateShortcutBoundary {
	return &CreateShortcutInteractor{
		factory: f,
		repo:    repo,
	}
}

func (i *CreateShortcutInteractor) CreateShortcut(ctx context.Context, input *CreateShortcutInput) (*CreateShortcutOutput, error) {
	lurl, err := url.Parse(input.LongURL)
	if err != nil {
		return nil, fmt.Errorf("can not parse as URL %w", err)
	}
	s, err := i.factory.NewShortcut(lurl)
	if err != nil {
		return nil, fmt.Errorf("can not create shortcut %w", err)
	}
	if err := i.repo.Save(ctx, s); err != nil {
		return nil, fmt.Errorf("can not save shortcut %w", err)
	}
	output := &CreateShortcutOutput{
		ShortcutID: s.ID(),
		LongURL:    s.LongURL().String(),
	}
	return output, nil
}
