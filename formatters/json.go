package formatters

import (
	"encoding/json"
)

func formatJSON(tree []Node) (string, error) {
	wrapper := map[string]any{
		"diff": tree,
	}

	data, err := json.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}
