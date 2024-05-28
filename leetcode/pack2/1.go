package main

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	var ans []string
	// проход по часам
	for h := 0; h < 12; h++ {
		// количество единиц в двочиной форме числа
		i := bits.OnesCount(uint(h))
		// проход по минутам
		for m := 0; m < 60; m++ {
			// количество единиц в двочиной форме числа
			j := bits.OnesCount(uint(m))
			// если сумма единиц совпадет таргету - выводим
			if turnedOn == i+j {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}
