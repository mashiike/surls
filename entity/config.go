package entity

type Config struct {
	//AvailableLetters For ShortcutID
	AvailableLetters string `json:"available_letters,omitempty"`
	//MinimumLength is the minimum length of the ShortcutID
	MinimumLength int `json:"minimum_length,omitempty"`
}

const (
	LowerAlphabets = "abcdefghijklmnopqrstuvwxyz"
	UpperAlphabets = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumberLetters  = "0123456789"
)

//NewDefaultConfig fills in the defaults and creates *Config
func NewDefaultConfig() *Config {
	return &Config{
		AvailableLetters: LowerAlphabets + UpperAlphabets + NumberLetters,
		MinimumLength:    5,
	}
}
