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
	s.root = remove(s.root, str)
	s.size--
	if s.size < 0 {
		s.size = 0
	}
}

func (s *Set) Contains(str string) bool {
	node := get(s.root, str)
	return node != nil
}

func (s *Set) Size() int64 {
	return s.size
}

func (s *Set) ForEach(action func(string)) {
	traverse(s.root, func(s string, _ ...interface{}) {
		action(s)
	})
}

func (s *Set) ForEachWithPrefix(prefix string, action func(string)) {
	node := getWithPrefix(s.root, prefix)
	traverse(node, func(s string, _ ...interface{}) {
		action(s)
	})
}
