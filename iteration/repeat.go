package iteration

import ()

const repeatCount = 5

func Repeat(character string) (result string) {
	for i := 0; i < repeatCount; i++ {
		result += character
	}
	return
}
