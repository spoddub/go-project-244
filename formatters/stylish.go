package formatters

import (
	"fmt"
	"sort"
	"strings"
)

const indentSize = 4

func formatStylish(tree []Node) string {
	var b strings.Builder
	b.WriteString("{\n")
	formatStylishNodes(&b, tree, 1)
	b.WriteString("}")
	return b.String()
}

func formatStylishNodes(b *strings.Builder, nodes []Node, depth int) {
	indent := strings.Repeat(" ", depth*indentSize-2)

	for _, n := range nodes {
		switch n.Type {
		case nodeNested:
			fmt.Fprintf(b, "%s  %s: {\n", indent, n.Key)
			formatStylishNodes(b, n.Children, depth+1)
			closingIndent := strings.Repeat(" ", depth*indentSize)
			fmt.Fprintf(b, "%s}\n", closingIndent)

		case nodeUnchanged:
			valStr := stringifyStylish(n.Value, depth)
			fmt.Fprintf(b, "%s  %s: %s\n", indent, n.Key, valStr)

		case nodeAdded:
			valStr := stringifyStylish(n.Value, depth)
			fmt.Fprintf(b, "%s+ %s: %s\n", indent, n.Key, valStr)

		case nodeRemoved:
			valStr := stringifyStylish(n.Value, depth)
			fmt.Fprintf(b, "%s- %s: %s\n", indent, n.Key, valStr)

		case nodeUpdated:
			oldValStr := stringifyStylish(n.OldValue, depth)
			newValStr := stringifyStylish(n.NewValue, depth)
			fmt.Fprintf(b, "%s- %s: %s\n", indent, n.Key, oldValStr)
			fmt.Fprintf(b, "%s+ %s: %s\n", indent, n.Key, newValStr)
		}
	}
}

func stringifyStylish(value any, depth int) string {
	if value == nil {
		return "null"
	}

	if m, ok := toMap(value); ok {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var b strings.Builder
		b.WriteString("{\n")

		for _, key := range keys {
			val := m[key]
			indent := strings.Repeat(" ", (depth+1)*indentSize)
			fmt.Fprintf(&b, "%s%s: %s\n", indent, key, stringifyStylish(val, depth+1))
		}

		closingIndent := strings.Repeat(" ", depth*indentSize)
		fmt.Fprintf(&b, "%s}", closingIndent)

		return b.String()
	}

	return fmt.Sprintf("%v", value)
}
