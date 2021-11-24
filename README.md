[![Package Version](https://img.shields.io/github/v/release/jpholanda/radixtree?style=for-the-badge)](https://github.com/jpholanda/radixtree/releases)
[![Go Version](https://img.shields.io/github/go-mod/go-version/jpholanda/radixtree?style=for-the-badge)](https://golang.org/)
[![Code Coverage](https://img.shields.io/scrutinizer/coverage/g/jpholanda/radixtree?style=for-the-badge)](https://scrutinizer-ci.com/g/jpholanda/radixtree/?branch=master)

# radixtree
Radix tree implementation in Go.

Exports a Set and a Map structures using a radix tree as the underlying structure.

Supported operations are:
- Add/Remove: inserts/deletes words into/from the tree. Linear on the size of the word.   
- Contains: checks whether the tree has a given word. Linear on the size of the word.
- ForEach: executes a callback for each word in the tree. Linear on the size of the tree.
- ForEachWithPrefix: executes a callback for each work in the tree with the given prefix. Linear on the size of the prefix and the number of words in the tree with that prefix.