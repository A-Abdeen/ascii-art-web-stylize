package asciiart

import (
	"os"
)

func AsciiArt(inputString string, banner string) string {
	if inputString == "" {
		return "400: " + "Input text missing"
	}

	switch {
	case banner == "standard":
		banner = "standard.txt"
	case banner == "shadow":
		banner = "shadow.txt"
	case banner == "thinkertoy":
		banner = "thinkertoy.txt"
	case banner == "internal error":
		banner = "internal error.txt"
	default:
		return "400: " + "Banner use required: available formats are: standard, shadow or thinkertoy."
	}
	sourceFile, err := os.ReadFile(banner)
	if err != nil {
		return "500: " + err.Error()
	}
	// Main function: Splitting (split string based on newline position)
	// ∟--> Sub function: Formatting (change input to allow use of newline & qoutation marks)
	splitInput := LineSplitter(inputString, InputFormatter)

	// Main function: Printing (printing the row of characters within input string)
	// ∟--> Sub function: Parsing (parsing the data of the 8 rows to print sequentially)
	return RowPrinter(splitInput, sourceFile, RowParser)

}
