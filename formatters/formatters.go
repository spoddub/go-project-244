package formatters

import "fmt"

func Format(data1, data2 map[string]any, format string) (string, error) {
	tree := buildDiff(data1, data2)

	switch format {
	case "", "stylish":
		return formatStylish(tree), nil
	case "plain":
		return formatPlain(tree), nil
	default:
		return "", fmt.Errorf("unsupported format %q", format)
	}
}
