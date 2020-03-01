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
	m.root = remove(m.root, str)
	m.size--
	if m.size < 0 {
		m.size = 0
	}
}

func (m *Map) Get(str string) interface{} {
	node := get(m.root, str)
	if node == nil {
		return nil
	}
	return node.data
}

func (m *Map) Size() int64 {
	return m.size
}

func (m *Map) ForEach(action func(string, interface{})) {
	traverse(s.root, func(s string, optdata ...interface{}) {
		if len(optdata) > 0 {
			action(s, optdata[0])
		} else {
			action(s, nil)
		}
	})
}

func (m *Map) ForEachWithPrefix(prefix string, action func(string, interface{})) {
	node := getWithPrefix(m.root, prefix)
	traverse(node, func(s string, optdata ...interface{}) {
		if len(optdata) > 0 {
			action(s, optdata[0])
		} else {
			action(s, nil)
		}
	})
}
