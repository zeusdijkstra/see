package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Test Markdown File</title>
  </head>
  <body>`
	footer = `
  </body>
</html>
`
)

func main() {
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, out io.Writer) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)

	// create temp file and check for errors
	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}

	outName := temp.Name()

	fmt.Fprintln(out, outName)

	return saveHTML(outName, htmlData)
}

func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	var buffer bytes.Buffer

	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHTML(outFname string, data []byte) error {
	return os.WriteFile(outFname, data, 0644)
}
