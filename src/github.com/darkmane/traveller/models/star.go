package models

import (
	"encoding/json"
	"fmt"
	"math/bits"
)

const NEAR_ORBIT= -1
const FAR_ORBIT = 1<<(bits.UintSize -1) -1 // MAX_INT
const CENTER = -1 <<(bits.UintSize -1) // MIN_INT
type Star struct {
	Id int
	Class StellarClass `json:"class"`
	Size StellarSize `json:"size"`
	Orbit int `json:"orbit"`
}

type StellarClass int {
	B = iota
	A 
	M
	K
	G
	F
}

type StellarSize int {
	Ia = iota
	Ib
	II
	III
	IV
	V
	VI
	D
}

type StarPosition int {
	PRIMARY=iota
	SECONDARY
	TERTIARY
}

