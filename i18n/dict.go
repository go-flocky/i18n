package i18n

import (
	"fmt"
	"strings"
)

type Dictionary struct {
	Value     string
	ChildDict map[string]*Dictionary
}

func buildDictTree(data map[string]any) *Dictionary {
	root := &Dictionary{
		ChildDict: make(map[string]*Dictionary),
	}

	for k, v := range data {
		switch val := v.(type) {
		case string:
			root.ChildDict[k] = &Dictionary{Value: val}
		case map[string]any:
			root.ChildDict[k] = buildDictTree(val)
		default:
			root.ChildDict[k] = &Dictionary{Value: fmt.Sprint(val)}
		}
	}

	return root
}

func mergeDict(dst, src *Dictionary) {
	if dst.ChildDict == nil {
		dst.ChildDict = make(map[string]*Dictionary)
	}

	for key, srcVal := range src.ChildDict {
		if dstVal, exists := dst.ChildDict[key]; exists {
			mergeDict(dstVal, srcVal)
		} else {
			dst.ChildDict[key] = srcVal
		}
	}
}

func (d *Dictionary) PrintTree() string {
	return d.printTreeHelper(0)
}

func (d *Dictionary) printTreeHelper(level int) string {
	result := ""
	indent := strings.Repeat("  ", level)

	for key, child := range d.ChildDict {
		if child.Value != "" && len(child.ChildDict) == 0 {
			result += fmt.Sprintf("%s%s: %s\n", indent, key, child.Value)
		} else {
			result += fmt.Sprintf("%s%s:\n", indent, key)
			result += child.printTreeHelper(level + 1)
		}
	}
	return result
}
