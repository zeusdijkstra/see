package main

import (
	"bytes"
	"os"
	"regexp"
	"strings"
	"testing"
)

const (
	inputFile  = "./testdata/test_inputFile.md"
	goldenFile = "./testdata/test_goldenFile.md.html"
)

func normalizeHTML(s string) string {
	reBetweenTags := regexp.MustCompile(`>\s+<`)
	s = reBetweenTags.ReplaceAllString(s, "><")

	reSpaces := regexp.MustCompile(`\s{2,}`)
	s = reSpaces.ReplaceAllString(s, " ")

	reCodeNewline := regexp.MustCompile(`(?s)<code>(.*?)\s*</code>`)
	s = reCodeNewline.ReplaceAllString(s, "<code>$1</code>")

	s = strings.TrimSpace(s)
	return s
}

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	result, err := parseContent(input, "")
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	normalizedResult := normalizeHTML(string(result))
	normalizedExpected := normalizeHTML(string(expected))

	if normalizedResult != normalizedExpected {
		t.Logf("golden (normalized): \n%s\n", normalizedExpected)
		t.Logf("result (normalized): \n%s\n", normalizedResult)
		t.Error("Result content does not match golden file (after normalization)")
	}
}

func TestRun(t *testing.T) {
	var mockStdOut bytes.Buffer

	if err := run(inputFile, "", &mockStdOut, true); err != nil {
		t.Fatal(err)
	}

	// get the value out of the buffer
	resultFile := strings.TrimSpace(mockStdOut.String())

	result, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	normalizedResult := normalizeHTML(string(result))
	normalizedExpected := normalizeHTML(string(expected))

	if normalizedResult != normalizedExpected {
		t.Logf("golden (normalized): \n%s\n", normalizedExpected)
		t.Logf("result (normalized): \n%s\n", normalizedResult)
		t.Error("Result content does not match golden file (after normalization)")
	}
	os.Remove(resultFile)
}
