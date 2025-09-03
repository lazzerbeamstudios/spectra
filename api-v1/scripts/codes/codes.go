package codes

import "math/rand"

func GenerateRandomLetters() string {
	validLetters := "abcdefghjkmnopqrstuvwxyz"
	result := ""
	for i := 0; i < 3; i++ {
		index := rand.Intn(len(validLetters))
		result += string(validLetters[index])
	}
	return result
}
