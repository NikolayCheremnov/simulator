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

func (i *Intgen) IntN(max int) int {
	return i.Int(0, max)
}

func IntRandom() *Intgen {
	rand.Seed(time.Now().UnixNano())
	return &Intgen{}
}
