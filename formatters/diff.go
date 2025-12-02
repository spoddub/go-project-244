// formatters/diff.go
package formatters

import (
	"reflect"
	"sort"
)

type nodeType string

const (
	nodeAdded     nodeType = "added"
	nodeRemoved   nodeType = "deleted"
	nodeUnchanged nodeType = "unchanged"
	nodeUpdated   nodeType = "changed"
	nodeNested    nodeType = "nest"
)

type Node struct {
	Key      string   `json:"key"`
	Type     nodeType `json:"type"`
	Value    any      `json:"value,omitempty"`
	OldValue any      `json:"oldValue,omitempty"`
	NewValue any      `json:"newValue,omitempty"`
	Children []Node   `json:"children,omitempty"`
}

func buildDiff(data1, data2 map[string]any) []Node {
	keysSet := make(map[string]struct{})

	for k := range data1 {
		keysSet[k] = struct{}{}
	}
	for k := range data2 {
		keysSet[k] = struct{}{}
	}

	keys := make([]string, 0, len(keysSet))
	for k := range keysSet {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := make([]Node, 0, len(keys))

	for _, key := range keys {
		v1, ok1 := data1[key]
		v2, ok2 := data2[key]

		switch {
		case ok1 && ok2:
			m1, m1ok := toMap(v1)
			m2, m2ok := toMap(v2)

			if m1ok && m2ok {
				children := buildDiff(m1, m2)
				result = append(result, Node{
					Key:      key,
					Type:     nodeNested,
					Children: children,
				})
			} else if reflect.DeepEqual(v1, v2) {
				result = append(result, Node{
					Key:   key,
					Type:  nodeUnchanged,
					Value: v1,
				})
			} else {
				result = append(result, Node{
					Key:      key,
					Type:     nodeUpdated,
					OldValue: v1,
					NewValue: v2,
				})
			}

		case ok1 && !ok2:
			result = append(result, Node{
				Key:   key,
				Type:  nodeRemoved,
				Value: v1,
			})

		case !ok1 && ok2:
			result = append(result, Node{
				Key:   key,
				Type:  nodeAdded,
				Value: v2,
			})
		}
	}

	return result
}

func toMap(v any) (map[string]any, bool) {
	if v == nil {
		return nil, false
	}

	if m, ok := v.(map[string]any); ok {
		return m, true
	}

	if m, ok := v.(map[string]interface{}); ok {
		res := make(map[string]any, len(m))
		for k, vv := range m {
			res[k] = vv
		}
		return res, true
	}

	return nil, false
}
