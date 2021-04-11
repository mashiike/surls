package randstr

import (
	"math"
	"math/rand"
	"time"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

type Generator struct {
	letters       string
	letterIdxBits int64
	letterIdxMask int64
	letterIdxMax  int64
}

func New(letters string) *Generator {
	g := &Generator{
		letters: letters,
	}
	g.letterIdxBits = int64(math.Ceil(math.Log2(float64(len(letters)))))
	g.letterIdxMask = 1<<g.letterIdxBits - 1
	g.letterIdxMax = 63 / g.letterIdxBits
	return g
}

func (g *Generator) Generate(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), g.letterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), g.letterIdxMax
		}
		idx := int(cache & g.letterIdxMask)
		if idx < len(g.letters) {
			b[i] = g.letters[idx]
			i--
		}
		cache >>= g.letterIdxBits
		remain--
	}
	return string(b)
}

func (g *Generator) Letters() string {
	return g.letters
}
