package code

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fixturePath(filename string) string {
	return filepath.Join("testdata", "fixture", filename)
}

func TestGenDiffNestedJSONStylish(t *testing.T) {
	file1 := fixturePath("file1.json")
	file2 := fixturePath("file2.json")
	expectedPath := fixturePath("nested_stylish.txt")

	got, err := GenDiff(file1, file2, "stylish")
	if !assert.NoError(t, err, "GenDiff returned error for json stylish") {
		return
	}

	expectedBytes, err := os.ReadFile(expectedPath)
	if !assert.NoError(t, err, "cannot read expected stylish fixture") {
		return
	}

	expected := strings.TrimSpace(string(expectedBytes))
	got = strings.TrimSpace(got)

	assert.Equal(t, expected, got)
}

func TestGenDiffNestedYAMLStylish(t *testing.T) {
	file1 := fixturePath("file1.yml")
	file2 := fixturePath("file2.yml")
	expectedPath := fixturePath("nested_stylish.txt")

	got, err := GenDiff(file1, file2, "stylish")
	if !assert.NoError(t, err, "GenDiff returned error for yaml stylish") {
		return
	}

	expectedBytes, err := os.ReadFile(expectedPath)
	if !assert.NoError(t, err, "cannot read expected stylish fixture") {
		return
	}

	expected := strings.TrimSpace(string(expectedBytes))
	got = strings.TrimSpace(got)

	assert.Equal(t, expected, got)
}

func TestGenDiffNestedJSONPlain(t *testing.T) {
	file1 := fixturePath("file1.json")
	file2 := fixturePath("file2.json")
	expectedPath := fixturePath("nested_plain.txt")

	got, err := GenDiff(file1, file2, "plain")
	if !assert.NoError(t, err, "GenDiff returned error for json plain") {
		return
	}

	expectedBytes, err := os.ReadFile(expectedPath)
	if !assert.NoError(t, err, "cannot read expected plain fixture") {
		return
	}

	expected := strings.TrimSpace(string(expectedBytes))
	got = strings.TrimSpace(got)

	assert.Equal(t, expected, got)
}

func TestGenDiffNestedYAMLPlain(t *testing.T) {
	file1 := fixturePath("file1.yml")
	file2 := fixturePath("file2.yml")
	expectedPath := fixturePath("nested_plain.txt")

	got, err := GenDiff(file1, file2, "plain")
	if !assert.NoError(t, err, "GenDiff returned error for yaml plain") {
		return
	}

	expectedBytes, err := os.ReadFile(expectedPath)
	if !assert.NoError(t, err, "cannot read expected plain fixture") {
		return
	}

	expected := strings.TrimSpace(string(expectedBytes))
	got = strings.TrimSpace(got)

	assert.Equal(t, expected, got)
}
