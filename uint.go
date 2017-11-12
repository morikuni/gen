package dummy

import (
	"math"
	"math/rand"
	"strconv"
)

// Uint returns a uint generator.
func Uint(r *rand.Rand) *UintGen {
	return &UintGen{
		r,
		math.MaxUint64,
		0,
	}
}

// UintGen is a uint generator.
type UintGen struct {
	r   *rand.Rand
	max uint64
	min uint64
}

// Max sets the maximum value of a random uint value.
func (g *UintGen) Max(n uint) *UintGen {
	return g.Max64(uint64(n))
}

// Min sets the minimum value of a random uint value.
func (g *UintGen) Min(n uint) *UintGen {
	return g.Min64(uint64(n))
}

// Max sets the maximum value of a random uint value with uint64.
func (g *UintGen) Max64(n uint64) *UintGen {
	g.max = n
	return g
}

// Min sets the minimum value of a random uint value with uint64.
func (g *UintGen) Min64(n uint64) *UintGen {
	g.min = n
	return g
}

// Gen returns a uint value.
func (g *UintGen) Gen() uint {
	return uint(g.Gen64())
}

// GenP returns a uint pointer.
func (g *UintGen) GenP() *uint {
	u := g.Gen()
	return &u
}

// Gen64P returns a uint64 pointer.
func (g *UintGen) Gen64P() *uint64 {
	u := g.Gen64()
	return &u
}

// GenString returns a uint64 value as a string value.
func (g *UintGen) GenString() string {
	return strconv.FormatUint(g.Gen64(), 10)
}

// GenString returns a uint64 value as a string value.
func (g *UintGen) GenStringP() *string {
	s := g.GenString()
	return &s
}
