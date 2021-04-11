package entity

import (
	"net/url"
	"strings"
)

type ShortcutID string

func (id ShortcutID) AsURL(baseURL *url.URL) *url.URL {
	cloned := *baseURL
	cloned.Path = strings.TrimRight(cloned.Path, "/") + "/" + string(id)
	return &cloned
}

type Shortcut struct {
	id      ShortcutID
	longURL *url.URL
}

func NewShortcut(id ShortcutID, longURL *url.URL) *Shortcut {
	cloned := *longURL
	return &Shortcut{
		id:      id,
		longURL: &cloned,
	}
}

func (s *Shortcut) ID() ShortcutID {
	return s.id
}

func (s *Shortcut) LongURL() *url.URL {
	cloned := *s.longURL
	return &cloned
}
