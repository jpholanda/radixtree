package radixtree

type Map struct {
	root *radixNode
	size int64
}

func (m *Map) Add(str string, data interface{}) {
	m.root = add(m.root, str, data)
	m.size++
}

func (m *Map) Remove(str string) {
	var removed bool
	m.root, removed = remove(m.root, str)
	if removed {
		m.size--
	}
}

func (m *Map) Get(str string) (interface{}, bool) {
	node := get(m.root, str)
	if node == nil || !node.final {
		return nil, false
	}
	return node.data, true
}

func (m *Map) Size() int64 {
	return m.size
}

func (m *Map) ForEach(action func(string, interface{})) {
	traverse(m.root, action)
}

func (m *Map) ForEachWithPrefix(prefix string, action func(string, interface{})) {
	node, buffer := getWithPrefix(m.root, prefix)
	traverseRecursive(node, buffer, action)
}
