package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filename := flag.String("file", "", "Markdown file to preview")
	tFname := flag.String("t", "", "Template name or file path")
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

func run(filename, tFname string, out io.Writer, skipPreview bool) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	htmlData, err := parseContent(input, tFname, filename)
	if err != nil {
		return fmt.Errorf("parse content: %w", err)
	}

	if skipPreview {
		base := filepath.Base(filename)
		name := strings.TrimSuffix(base, filepath.Ext(base)) + ".html"

		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("getwd: %w", err)
		}

		outPath := filepath.Join(cwd, name)

		if err := saveHTML(outPath, htmlData); err != nil {
			return fmt.Errorf("save html: %w", err)
		}

		fmt.Fprintln(out, outPath)
		return nil
	}

	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	defer temp.Close()

	outName := temp.Name()
	defer os.Remove(outName)

	fmt.Fprintln(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return fmt.Errorf("save html: %w", err)
	}

	if err := preview(outName); err != nil {
		return fmt.Errorf("preview: %w", err)
	}
	return nil
}
