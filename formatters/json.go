package formatters

import (
	"encoding/json"
)

func formatJSON(tree []Node) (string, error) {
	data, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
