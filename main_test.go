package main

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	resultFile = "test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)

// normalizeHTML removes insignificant differences in whitespace, indentation, and newlines
func normalizeHTML(s string) string {
	// collapse spaces and newlines between tags
	reBetweenTags := regexp.MustCompile(`>\s+<`)
	s = reBetweenTags.ReplaceAllString(s, "><")

	// remove extra spaces, tabs, and newlines
	reSpaces := regexp.MustCompile(`\s{2,}`)
	s = reSpaces.ReplaceAllString(s, " ")

	// remove newlines before </code>
	reCodeNewline := regexp.MustCompile(`(?s)<code>(.*?)\s*</code>`)
	s = reCodeNewline.ReplaceAllString(s, "<code>$1</code>")

	// trim overall whitespace
	s = strings.TrimSpace(s)
	return s
}

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	result := parseContent(input)
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
	if err := run(inputFile); err != nil {
		t.Fatal(err)
	}

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
}
