package kvtree

import (
	"strings"
)

type TreeNode[T any] struct {
	Value    *T
	Children map[string]*TreeNode[T]
}

type KeyValueTree[T any] struct {
	Root      *TreeNode[T]
	Separator string
}

func NewTree[T any](keySeparator string) *KeyValueTree[T] {
	return &KeyValueTree[T]{
		Separator: keySeparator,
		Root: &TreeNode[T]{
			Children: make(map[string]*TreeNode[T]),
		},
	}
}

func (t *KeyValueTree[T]) Set(key string, value T) {
	if key == "" {
		return
	}

	node := t.Root
	for _, part := range strings.Split(key, t.Separator) {
		if part == "" {
			continue
		}
		if node.Children[part] == nil {
			node.Children[part] = &TreeNode[T]{
				Children: make(map[string]*TreeNode[T]),
			}
		}
		node = node.Children[part]
	}

	node.Value = &value
}

func (t *KeyValueTree[T]) Get(key string) (*T, bool) {
	if key == "" {
		return nil, false
	}

	node := t.Root
	for _, part := range strings.Split(key, t.Separator) {
		if part == "" {
			continue
		}
		child, ok := node.Children[part]
		if !ok {
			return nil, false
		}
		node = child
	}

	if node.Value == nil {
		return nil, false
	}

	return node.Value, true
}
