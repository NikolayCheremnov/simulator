package random

import (
	"math/rand"
	"time"
)

type Intgen struct {
}

func (i *Intgen) Int(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func IntRandom() *Intgen {
	rand.Seed(time.Now().UnixNano())
	return &Intgen{}
}
