package asciiart

func RowParser(sourceFile []byte, rowLocation int) []byte {
	charRowData := []byte{}
	count := 1
	lineNumber := 1 
	for i := 0; i >= 0; i++ {
		if sourceFile[i] == 10 {   // if new line found increase count
			count++
		}
		if count == rowLocation {
			lineNumber = i + 1
			break
		}
	}
	for ; lineNumber > 0; lineNumber++ {
		if sourceFile[lineNumber] == 10 { // if new line found stop appending char data
			break
		}
		charRowData = append(charRowData, sourceFile[lineNumber])
	}
	return charRowData
}
