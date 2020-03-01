package radixtree

import "bytes"

type radixNode struct {
	children map[byte]*radixNode
	part     string
	final    bool
	data     interface{}
}

func newRadixNode(part string, final bool, optdata ...interface{}) *radixNode {
	data := interface{}(nil)
	if len(optdata) > 0 {
		data = optdata[0]
	}

	return &radixNode{
		part:     part,
		final:    final,
		children: make(map[byte]*radixNode),
		data:     data,
	}
}

func (root *radixNode) addChild(child *radixNode) {
	root.children[child.part[0]] = child
}

func (root *radixNode) removeChild(child *radixNode) {
	delete(root.children, child.part[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func commonPrefixLength(a, b string) int {
	minlen := min(len(a), len(b))
	for i := 0; i < minlen; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return minlen
}

func add(root *radixNode, str string, optdata ...interface{}) *radixNode {
	data := interface{}(nil)
	if len(optdata) > 0 {
		data = optdata[0]
	}

	if root == nil {
		return newRadixNode(str, true, data)
	}

	lenPrefix := commonPrefixLength(root.part, str)

	matchExactly := lenPrefix == len(root.part) && lenPrefix == len(str)
	if matchExactly {
		root.final = true
		return root
	}

	rootIsPrefixOfString := lenPrefix == len(root.part)
	if rootIsPrefixOfString {
		endStr := str[lenPrefix:]

		var newChild *radixNode

		candidateChild, exists := root.children[endStr[0]]
		if exists {
			newChild = add(candidateChild, endStr, data)
		} else {
			newChild = newRadixNode(endStr, true, data)
		}

		root.addChild(newChild)
		return root
	}

	// if we got here, then the common prefix must be split
	// into a separate node, which will be the new root

	var newRoot *radixNode
	newStringIsPrefixOfRoot := lenPrefix == len(str)
	if newStringIsPrefixOfRoot {
		prefix := str
		newRoot = newRadixNode(prefix, true, data)
	} else {
		prefix := root.part[:lenPrefix]
		newRoot = newRadixNode(prefix, false)

		// newRoot will have two children, one with the
		// final part of the old root, and the other with
		// the final part of the new string
		endStr := str[lenPrefix:]
		newChild := newRadixNode(endStr, true, data)
		newRoot.addChild(newChild)
	}

	// reuse root to avoid setting up the children again
	oldRoot := root
	oldRoot.part = root.part[lenPrefix:]

	newRoot.addChild(oldRoot)
	return newRoot
}

func remove(root *radixNode, str string) *radixNode {
	if root == nil {
		return nil
	}

	lenPrefix := commonPrefixLength(root.part, str)

	matchExactly := lenPrefix == len(root.part) && lenPrefix == len(str)
	if matchExactly {
		if !root.final {
			return root
		}

		rootIsWord := true

		hasNoChildren := len(root.children) == 0
		if rootIsWord && hasNoChildren {
			return nil
		}

		hasOneChild := len(root.children) == 1
		if rootIsWord && hasOneChild {
			mergeWithSingleChild(root)
			return root
		}

		// if we got here, root has children, so we can just unset the final flag

		root.final = false
		root.data = nil
		return root
	}

	rootIsPrefixOfString := lenPrefix == len(root.part)
	if rootIsPrefixOfString {
		endStr := str[lenPrefix:]

		childCandidate, exists := root.children[endStr[0]]
		if !exists {
			return root
		}

		shouldRemoveChild := remove(childCandidate, endStr) == nil
		if shouldRemoveChild {
			root.removeChild(childCandidate)
			if !root.final && len(root.children) == 1 {
				mergeWithSingleChild(root)
			}
		}
	}

	return root
}

func mergeWithSingleChild(node *radixNode) {
	// use range for to select the only child
	var child *radixNode
	for _, child = range node.children {
	}

	node.part += child.part
	node.children = child.children
	node.final = child.final
	node.data = child.data
}

func get(root *radixNode, str string) *radixNode {
	if root == nil {
		return nil
	}

	lenPrefix := commonPrefixLength(root.part, str)

	matchExactly := lenPrefix == len(root.part) && lenPrefix == len(str)
	if matchExactly {
		return root
	}

	rootIsPrefixOfString := lenPrefix == len(root.part)
	if rootIsPrefixOfString {
		endStr := str[lenPrefix:]

		childCandidate, exists := root.children[endStr[0]]
		if !exists {
			return nil
		}

		return get(childCandidate, endStr)
	}

	return nil
}

func getWithPrefix(root *radixNode, pattern string) *radixNode {
	if root == nil {
		return nil
	}

	lenPrefix := commonPrefixLength(root.part, pattern)

	patternIsPrefixOrEqualToRoot := lenPrefix == len(pattern)
	if patternIsPrefixOrEqualToRoot {
		return root
	}

	rootIsPrefixOfPattern := lenPrefix == len(root.part)
	if rootIsPrefixOfPattern {
		endPattern := pattern[lenPrefix:]

		childCandidate, exists := root.children[endPattern[0]]
		if !exists {
			return nil
		}

		return getWithPrefix(childCandidate, endPattern)
	}

	return nil
}

func traverse(root *radixNode, action func(string, ...interface{})) {
	buffer := &bytes.Buffer{}
	traverseRecursive(root, buffer, action)
}

func traverseRecursive(root *radixNode, buffer *bytes.Buffer, action func(string, ...interface{})) {
	if root == nil {
		return
	}

	sizebefore := buffer.Len()

	_, _ = buffer.WriteString(root.part)
	if root.final {
		if root.data != nil {
			action(buffer.String(), root.data)
		} else {
			action(buffer.String())
		}
	}

	for _, child := range root.children {
		traverseRecursive(child, buffer, action)
	}
	buffer.Truncate(sizebefore)
}
