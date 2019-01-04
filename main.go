package main

import (
	"fmt"
	"large-numbers-law/subject"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	subjects := []subject.Interface{&subject.Coin{}}

	contexts := []subject.Context{
		{10, 100, r},
		{100, 100, r},
		{1000, 100, r},
		{10000, 100, r},
		{100000, 100, r},
		{1000000, 100, r},
		{10000000, 100, r},

		{100, 10, r},
		{100, 100, r},
		{100, 1000, r},
		{100, 10000, r},
		{100, 100000, r},
		{100, 1000000, r},
	}

	for _, context := range contexts {
		fmt.Printf("for %d groups and %d attempts:\n", context.Paricipants, context.Attempts)

		for _, subj := range subjects {
			subj.SetAttemptsNumber(context.Attempts)
			subj.SetParicipantsNumber(context.Paricipants)
			subj.SetRand(context.Rand)

			result := subj.Check()
			fmt.Printf("%s: average value: %f, average percent: %.2f%% \n", reflect.TypeOf(subj), result.AverageValue, result.AveragePercent*100)
		}

		fmt.Println()
	}
}
