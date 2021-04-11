package randstr_test

import (
	"regexp"
	"testing"

	"github.com/mashiike/surls/internal/randstr"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var validator = regexp.MustCompile("^[" + letters + "]*$")

func TestGenerator(t *testing.T) {
	g := randstr.New(letters)
	str := g.Generate(5)
	if len(str) != 5 {
		t.Errorf("random string length unexpected = %d", len(str))
	}
	if !validator.Match([]byte(str)) {
		t.Errorf("unexpected random string, validator not match. str is %s", str)
	}
}

func BenchmarkGenerator(b *testing.B) {
	g := randstr.New(letters)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Generate(5)
	}
}
