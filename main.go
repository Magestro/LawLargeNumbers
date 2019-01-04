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
	}

	for _, context := range contexts {
		fmt.Printf("for %d groups and %d attempts:\n", context.Groups, context.Attempts)

		for _, subj := range subjects {
			subj.SetAttemptsNumber(context.Attempts)
			subj.SetGroupsNumber(context.Groups)
			subj.SetRand(context.Rand)

			result := subj.Check()
			fmt.Printf("%s: average value: %f, average percent: %.2f%% \n", reflect.TypeOf(subj), result.AverageValue, result.AveragePercent*100)
		}

		fmt.Println()
	}
}
