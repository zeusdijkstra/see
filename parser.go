package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

//go:embed templates/*.html
var templateFS embed.FS

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)
	for _, name := range []string{"default", "minimal", "dark"} {
		t, err := template.ParseFS(templateFS, "templates/"+name+".html")
		if err != nil {
			panic(fmt.Sprintf("failed to parse template %s: %v", name, err))
		}
		templates[name] = t
	}
}

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

	if tFname == "" {
		t = templates["default"]
	} else if tmpl, ok := templates[tFname]; ok {
		t = tmpl
	} else {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, fmt.Errorf("parse custom template %q: %w", tFname, err)
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
