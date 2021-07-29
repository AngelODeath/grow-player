package fynelibs

import (
	rando "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func GenerateRandomNumber(guessMax int) int {
	const guessMin = 1
	guessResult := rand.Intn(guessMax-guessMin) + guessMin
	return guessResult
}

func GenerateRandomString(resultLength int) string {
		result := make([]byte, resultLength)
		_, err := rando.Read(result)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		for i := 0; i < resultLength; i++ {
			result[i] &= 0x7F
			for result[i] < 31 || result[i] == 127 {
				_, err = rando.Read(result[i : i+1])
				if err != nil {
					fmt.Println(err)
					return ""
				}
				result[i] &= 0x7F
			}
		}
		return string(result)
	}