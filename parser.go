package main

import (
	"bytes"
	"fmt"
	"html/template"

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
    <header>
        <small>Viewing file: {{ .FileName }}</small>
    </header>
    {{ .Body }}
</body>
</html>
`
)

type content struct {
	Title    string
	Body     template.HTML
	FileName string
}

func parseContent(input []byte, tFname string, filename string) ([]byte, error) {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	var t *template.Template
	var err error

	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, fmt.Errorf("parse custom template %q: %w", tFname, err)
		}
	} else {
		// otherwise fall back to the default template
		t, err = template.New("default").Parse(defaultTemplate)
		if err != nil {
			return nil, fmt.Errorf("parse default template: %w", err)
		}
	}

	c := content{
		Title:    "Test Markdown File",
		Body:     template.HTML(body),
		FileName: filename,
	}

	var buffer bytes.Buffer

	if err := t.Execute(&buffer, c); err != nil {
		return nil, fmt.Errorf("execute template: %w", err)
	}

	return buffer.Bytes(), nil
}
