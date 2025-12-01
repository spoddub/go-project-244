package code

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func fixturePath(filename string) string {
	return filepath.Join("testdata", "fixture", filename)
}

func TestGenDiffFlatJSONStylish(t *testing.T) {
	file1 := fixturePath("file1.json")
	file2 := fixturePath("file2.json")
	expectedPath := fixturePath("flat_stylish.txt")

	got, err := GenDiff(file1, file2, "stylish")
	if err != nil {
		t.Fatalf("GenDiff returned error: %v", err)
	}

	expectedBytes, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("cannot read expected fixture: %v", err)
	}

	expected := strings.TrimSpace(string(expectedBytes))
	got = strings.TrimSpace(got)

	if got != expected {
		t.Fatalf("unexpected diff result\nGot:\n%q\n\nExpected:\n%q", got, expected)
	}
}
