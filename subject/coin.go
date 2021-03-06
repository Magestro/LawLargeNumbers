package subject

import (
	"math"
	"math/rand"
)

const (
	coinExceptedValue = (1 + -1) / 2
)

type Coin struct {
	attempts int64
	paricipants  int64
	rand     *rand.Rand
}

func (c *Coin) Check() (result CheckResult) {
	result.ExceptedValue = coinExceptedValue

	for i := int64(0); i < c.GetParicipantsNumber(); i++ {
		total := int64(0)
		for j := int64(0); j < c.GetAttemptsNumber(); j++ {
			if c.rand.Intn(2) > 0 {
				total++
			} else {
				total--
			}
		}

		total = int64(math.Abs(float64(total)))
		result.RealValues = append(result.RealValues, float64(total))
		result.RealPercents = append(result.RealPercents, float64(total)/float64(c.GetAttemptsNumber()))
	}

	var total float64 = 0
	for _, value := range result.RealValues {
		total += value
	}
	result.AverageValue = total / float64(c.GetParicipantsNumber())

	total = 0
	for _, value := range result.RealPercents {
		total += value
	}
	result.AveragePercent = total / float64(c.GetParicipantsNumber())

	return
}

func (c *Coin) SetParicipantsNumber(n int64) {
	c.paricipants = n
}

func (c *Coin) GetParicipantsNumber() int64 {
	return c.paricipants
}

func (c *Coin) SetRand(r *rand.Rand) {
	c.rand = r
}

func (c *Coin) SetAttemptsNumber(n int64) {
	c.attempts = n
}

func (c *Coin) GetAttemptsNumber() int64 {
	return c.attempts
}
