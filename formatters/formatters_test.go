package formatters

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"code/parsers"

	"github.com/stretchr/testify/assert"
)

func fixturePath(filename string) string {
	return filepath.Join("..", "testdata", "fixture", filename)
}

func TestFormatStylishNested(t *testing.T) {
	file1 := fixturePath("file1.json")
	file2 := fixturePath("file2.json")
	expectedPath := fixturePath("nested_stylish.txt")

	data1, err := parsers.Parse(file1)
	if !assert.NoError(t, err, "Parse should not return error for file1.json") {
		return
	}

	data2, err := parsers.Parse(file2)
	if !assert.NoError(t, err, "Parse should not return error for file2.json") {
		return
	}

	got, err := Format(data1, data2, "stylish")
	if !assert.NoError(t, err, "Format (stylish) returned error") {
		return
	}

	expectedBytes, err := os.ReadFile(expectedPath)
	if !assert.NoError(t, err, "cannot read stylish fixture") {
		return
	}

	expected := strings.TrimSpace(string(expectedBytes))
	got = strings.TrimSpace(got)

	assert.Equal(t, expected, got, "stylish formatter output mismatch")
}

func TestFormatPlainNested(t *testing.T) {
	file1 := fixturePath("file1.json")
	file2 := fixturePath("file2.json")
	expectedPath := fixturePath("nested_plain.txt")

	data1, err := parsers.Parse(file1)
	if !assert.NoError(t, err, "Parse should not return error for file1.json") {
		return
	}

	data2, err := parsers.Parse(file2)
	if !assert.NoError(t, err, "Parse should not return error for file2.json") {
		return
	}

	got, err := Format(data1, data2, "plain")
	if !assert.NoError(t, err, "Format (plain) returned error") {
		return
	}

	expectedBytes, err := os.ReadFile(expectedPath)
	if !assert.NoError(t, err, "cannot read plain fixture") {
		return
	}

	expected := strings.TrimSpace(string(expectedBytes))
	got = strings.TrimSpace(got)

	assert.Equal(t, expected, got, "plain formatter output mismatch")
}

func TestFormatJSONNested(t *testing.T) {
	file1 := fixturePath("file1.json")
	file2 := fixturePath("file2.json")

	data1, err := parsers.Parse(file1)
	if !assert.NoError(t, err, "Parse should not return error for file1.json") {
		return
	}

	data2, err := parsers.Parse(file2)
	if !assert.NoError(t, err, "Parse should not return error for file2.json") {
		return
	}

	got, err := Format(data1, data2, "json")
	if !assert.NoError(t, err, "Format (json) returned error") {
		return
	}

	var wrapper map[string]json.RawMessage
	if !assert.NoError(t, json.Unmarshal([]byte(got), &wrapper), "output must be valid JSON object") {
		return
	}

	rawDiff, ok := wrapper["diff"]
	if !assert.True(t, ok, "json must contain 'diff' field") {
		return
	}

	var nodes []Node
	if !assert.NoError(t, json.Unmarshal(rawDiff, &nodes), "diff field must be valid JSON array of nodes") {
		return
	}

	if assert.GreaterOrEqual(t, len(nodes), 4, "expected at least 4 top-level nodes") {
		var foundGroup2, foundGroup3 bool

		for _, n := range nodes {
			switch n.Key {
			case "group2":
				foundGroup2 = true
				assert.Equal(t, nodeRemoved, n.Type)
			case "group3":
				foundGroup3 = true
				assert.Equal(t, nodeAdded, n.Type)
			}
		}

		assert.True(t, foundGroup2, "node with key 'group2' must exist")
		assert.True(t, foundGroup3, "node with key 'group3' must exist")
	}
}

func TestFormatUnsupported(t *testing.T) {
	data1 := map[string]any{"key": "value1"}
	data2 := map[string]any{"key": "value2"}

	got, err := Format(data1, data2, "unknown-format")

	assert.Error(t, err, "expected error for unsupported format")
	assert.Equal(t, "", got, "output should be empty string on error")
}
