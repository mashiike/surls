package entity

import (
	"net/url"

	"github.com/mashiike/surls/internal/randstr"
)

type Factory struct {
	conf                *Config
	shortcutIDGenerator *randstr.Generator
}

func NewFactory(conf *Config) *Factory {
	return &Factory{
		conf:                conf,
		shortcutIDGenerator: randstr.New(conf.AvailableLetters),
	}
}

func (f *Factory) newShortcutID() (ShortcutID, error) {
	return ShortcutID(f.shortcutIDGenerator.Generate(f.conf.MinimumLength)), nil
}

func (f *Factory) NewShortcutID(idstr string) (ShortcutID, error) {
	return ShortcutID(idstr), nil
}

func (f *Factory) NewShortcut(longURL *url.URL) (*Shortcut, error) {
	id, err := f.newShortcutID()
	if err != nil {
		return nil, err
	}
	s := NewShortcut(
		id,
		longURL,
	)
	return s, nil
}
