package words

import (
	"math/rand"
	"time"
)

func Words() string {
	words := []string{
		"kota",
		"kimura",
		"ririko",
		"nakaoka",
		"tabetai",
		"yakiniku",
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(words))

	word := words[i]
	return word
}
