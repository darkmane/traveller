package util

import (
	"crypto/sha256"
	"math/rand"
)

// DiceGenerator allows for rolling dice
type DiceGenerator struct {
	rng *rand.Rand
}

// NewDiceGenerator Create new DiceGenerator
func NewDiceGenerator(seed string) DiceGenerator {
	h := sha256.New()
	h.Write([]byte(seed))

	bs := h.Sum(nil)
	var v int64 = 0
	for _, element := range bs {
		v += int64(element)
	}
	s := rand.NewSource(v)
	return DiceGenerator{rng: rand.New(s)}

}

// RollDiceWithModifier Roll an arbitrary number of d6 with a modifer at the end.
func (dg *DiceGenerator) RollDiceWithModifier(dice int, mod int) int {
	v := 0
	for index := 1; index <= dice; index++ {
		v += dg.rng.Intn(sides) + 1
	}
	return v + mod
}

// RollDice Roll a specified number of dice
func (dg *DiceGenerator) RollDice(dice int) int {
	return dg.RollDiceWithModifier(dice, 0)
}

// Roll 2d6
func (dg *DiceGenerator) Roll() int {
	return dg.RollDice(2)
}
