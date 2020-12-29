package iteration

// Repeat accepts a character and an integer and returns the character repeated
func Repeat(charInput string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += charInput
	}
	return repeated
}
