package radixtree

type Set struct {
	root *radixNode
	size int64
}

func (s *Set) Add(str string) {
	s.root = add(s.root, str)
	s.size++
}

func (s *Set) Remove(str string) {
	var removed bool
	s.root, removed = remove(s.root, str)
	if removed {
		s.size--
	}
}

func (s *Set) Contains(str string) bool {
	node := get(s.root, str)
	return node != nil && node.final
}

func (s *Set) Size() int64 {
	return s.size
}

func (s *Set) ForEach(action func(string)) {
	traverse(s.root, func(s string, _ interface{}) {
		action(s)
	})
}

func (s *Set) ForEachWithPrefix(prefix string, action func(string)) {
	node, buffer := getWithPrefix(s.root, prefix)
	traverseRecursive(node, buffer, func(s string, _ interface{}) {
		action(s)
	})
}
