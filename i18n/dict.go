package i18n

import "fmt"

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
