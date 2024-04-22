package faker

import (
	"math/rand"
	"time"
)

var faker = New("en")

type Faker struct {
	r      *rand.Rand
	locale string
}

func New(locale string) *Faker {
	return &Faker{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
