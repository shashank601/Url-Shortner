package shortcode

import "math/rand"


func Generate(input string) (string, error) {

	randomString := ""
	for i := 0; i < 6; i++ {
		randomString += string(rune(rand.Intn(26) + 'a'))
	}

	return randomString, nil
}
