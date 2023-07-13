package asciiart

func LineSplitter(rawInput string, f func(string) string) []string {
	formattedInput := f(rawInput)
	var oneLine string
	var linesArray []string
	for x, y := range formattedInput {
		if y == 10 {
			linesArray = append(linesArray, oneLine)
			oneLine = ""
		} else {
			oneLine = oneLine + string(y)
		}
		if x == len(formattedInput)-1 && oneLine != "" {
			linesArray = append(linesArray, oneLine)
		}
	}
	return linesArray

}
