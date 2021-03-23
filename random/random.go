package random

import (
	cryptorand "crypto/rand"
	"io"
	"math"
	"math/big"
	"math/rand"
	"sync"
	"time"
)

// Rand is a global *rand.Rand instance, which initialized with NewSource() source.
var Rand = rand.New(NewSource())

// Reader is a global, shared instance of a pseudorandom bytes generator.
// It doesn't consume entropy.
var Reader io.Reader = &reader{rnd: Rand}

func NewSource() rand.Source {
	var seed int64
	if cryptoseed, err := cryptorand.Int(cryptorand.Reader, big.NewInt(math.MaxInt64)); err != nil {
		// This should not happen, but worst-case fallback to time-based seed.
		seed = time.Now().UnixNano()
	} else {
		seed = cryptoseed.Int64()
	}
	return &lockedSource{
		src: rand.NewSource(seed),
	}
}

// copypaste from standard math/rand
type lockedSource struct {
	src rand.Source
	sync.Mutex
}

func (r *lockedSource) Int63() int64 {
	r.Lock()
	defer r.Unlock()
	return r.src.Int63()
}

func (r *lockedSource) Seed(seed int64) {
	r.Lock()
	defer r.Unlock()
	r.src.Seed(seed)
}

type reader struct {
	rnd *rand.Rand
}

func (r *reader) Read(b []byte) (int, error) {
	i := 0
	for {
		val := r.rnd.Int63()
		for val > 0 {
			b[i] = byte(val)
			i++
			if i == len(b) {
				return i, nil
			}
			val >>= 8
		}
	}
}
