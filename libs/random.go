package fynelibs

import (
	rando "crypto/rand"
	"math/rand"
)

func GenerateRandomNumber(guessMax int) int {
	const guessMin = 1
	guessResult := rand.Intn(guessMax-guessMin) + guessMin
	return guessResult
}

func GenerateRandomString(resultLength int) (string, error) {
		result := make([]byte, resultLength)
		_, err := rando.Read(result)
		if err != nil {
			return "", err
		}
		for i := 0; i < resultLength; i++ {
			result[i] &= 0x7F
			for result[i] < 31 || result[i] == 127 {
				_, err = rando.Read(result[i : i+1])
				if err != nil {
					return "", err
				}
				result[i] &= 0x7F
			}
		}
		return string(result), nil
	}