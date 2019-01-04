package subject

import (
	"math/rand"
)

type Context struct {
	Attempts int64
	Groups   int64
	Rand     *rand.Rand
}

type CheckResult struct {
	ExceptedValue  float64
	RealValues     []float64
	RealPercents   []float64
	AverageValue   float64
	AveragePercent float64
}

type Interface interface {
	SetAttemptsNumber(int64)

	GetAttemptsNumber() int64

	SetGroupsNumber(int64)

	GetGroupsNumber() int64

	SetRand(*rand.Rand)

	Check() CheckResult
}
