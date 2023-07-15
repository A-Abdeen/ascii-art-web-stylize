package asciiart


func IsAscii(input string) bool {
	for _, char := range input {
		if char == 10 || char == 13 {
			continue
		}
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}