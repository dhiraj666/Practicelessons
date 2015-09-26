package words

func CountWords(text string) int {
	count := 0
	for i := 0; i < len(text); i++ {
		count++
	}

	return count
}
