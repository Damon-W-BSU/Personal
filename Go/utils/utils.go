package utils

import "os"

// dumps bytes to a text file for viewing
func DumpToTxt(raw []byte) {
	f, err := os.Create("sorting/dump.txt")
	if err == nil {
		f.Write(raw)
	}
}
