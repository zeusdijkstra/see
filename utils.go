package main

import (
	"os"
)

func saveHTML(outFname string, data []byte) error {
	return os.WriteFile(outFname, data, 0644)
}
