package formatters

import (
	"fmt"
	"reflect"
	"strings"
)

func formatPlain(tree []Node) string {
	var lines []string
	buildPlainLines(tree, []string{}, &lines)
	return strings.Join(lines, "\n")
}

func buildPlainLines(nodes []Node, path []string, lines *[]string) {
	for _, n := range nodes {
		fullPath := append(path, n.Key)
		property := strings.Join(fullPath, ".")

		switch n.Type {
		case nodeNested:
			buildPlainLines(n.Children, fullPath, lines)

		case nodeAdded:
			valueStr := plainValue(n.Value)
			*lines = append(*lines,
				fmt.Sprintf("Property '%s' was added with value: %s", property, valueStr))

		case nodeRemoved:
			*lines = append(*lines,
				fmt.Sprintf("Property '%s' was removed", property))

		case nodeUpdated:
			fromStr := plainValue(n.OldValue)
			toStr := plainValue(n.NewValue)
			*lines = append(*lines,
				fmt.Sprintf("Property '%s' was updated. From %s to %s", property, fromStr, toStr))

		case nodeUnchanged:
			// ничего не выводим
		}
	}
}

func plainValue(v any) string {
	if isComplex(v) {
		return "[complex value]"
	}

	if v == nil {
		return "null"
	}

	switch vv := v.(type) {
	case string:
		return fmt.Sprintf("'%s'", vv)
	default:
		return fmt.Sprintf("%v", vv)
	}
}

func isComplex(v any) bool {
	if v == nil {
		return false
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array:
		return true
	default:
		return false
	}
}
