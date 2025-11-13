package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	defaultTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{{ .Title }}</title>
</head>
<body>
{{ .Body }}
</body>
</html>
`
)

type content struct {
	Title string
	Body  template.HTML
}

func main() {
	filename := flag.String("file", "", "Markdown file to preview")
	tFname := flag.String("t", "", "Alternate template name")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")

	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func preview(fname string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", fname)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", fname)
	case "darwin":
		cmd = exec.Command("open", fname)
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// use Start() so it doesn’t block while the viewer opens.
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start preview: %w", err)
	}

	// wait a short moment to help ensure the file opens properly
	time.Sleep(1 * time.Second)

	return nil
}

func run(filename, tFname string, out io.Writer, skipPreview bool) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	htmlData, err := parseContent(input, tFname)
	if err != nil {
		return fmt.Errorf("parse content: %w", err)
	}

	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	defer temp.Close()

	outName := temp.Name()
	fmt.Fprintln(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return fmt.Errorf("save html: %w", err)
	}

	if skipPreview {
		return nil
	}

	if err := preview(outName); err != nil {
		return fmt.Errorf("preview: %w", err)
	}

	// cleanup temp file asynchronously after giving the preview time to load
	// go func(name string) {
	// 	time.Sleep(10 * time.Second)
	// 	_ = os.Remove(name)
	// }(outName)
	// OR JUST DON'T REMOVE IT AT ALL...

	return nil
}

func parseContent(input []byte, tFname string) ([]byte, error) {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// If user does not provide a custom template
	t, err := template.New("newTemplate").Parse(defaultTemplate)
	if err != nil {
		return nil, fmt.Errorf("parse default template: %w", err)
	}

	// If user provides a custom template
	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, fmt.Errorf("parse custom template %q: %w", tFname, err)
		}
	}

	c := content{
		Title: "Test Markdown File",
		Body:  template.HTML(body),
	}

	var buffer bytes.Buffer

	if err := t.Execute(&buffer, c); err != nil {
		return nil, fmt.Errorf("execute template: %w", err)
	}

	return buffer.Bytes(), nil
}

func saveHTML(outFname string, data []byte) error {
	return os.WriteFile(outFname, data, 0644)
}
