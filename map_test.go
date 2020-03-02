package radixtree_test

import (
	"testing"

	"github.com/jpholanda/radixtree"
	"github.com/stretchr/testify/assert"
)

type pair struct {
	key  string
	data interface{}
}

func TestMapBasic(t *testing.T) {
	t.Parallel()

	t.Run("empty contains nothing", func(t *testing.T) {
		rmap := radixtree.Map{}

		_, exists := rmap.Get("aaa")
		assert.False(t, exists)
	})

	t.Run("contain after add", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("bbb", 222)

		data, exists := rmap.Get("bbb")

		assert.True(t, exists)
		assert.EqualValues(t, 222, data)
	})

	t.Run("not contain after add and remove", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("ccc", 333)
		rmap.Remove("ccc")

		_, exists := rmap.Get("ccc")

		assert.False(t, exists)
	})

	t.Run("size after add", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("ddd", 444)

		assert.EqualValues(t, 1, rmap.Size())
	})

	t.Run("size after remove", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Remove("eee")

		assert.EqualValues(t, 0, rmap.Size())
	})

	t.Run("size after add then remove", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("fff", 666)
		rmap.Remove("fff")

		assert.EqualValues(t, 0, rmap.Size())
	})

	t.Run("for each when empty", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.ForEach(func(_ string, _ interface{}) {
			panic("should not be executed")
		})
	})

	t.Run("for each when not empty", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("ggg", 777)
		rmap.Add("hhh", 888)
		rmap.Add("iii", 999)

		count := map[pair]int{}
		rmap.ForEach(func(s string, data interface{}) {
			count[pair{key: s, data: data}]++
		})

		assert.Len(t, count, 3)
		assert.EqualValues(t, 1, count[pair{key: "ggg", data: 777}])
		assert.EqualValues(t, 1, count[pair{key: "hhh", data: 888}])
		assert.EqualValues(t, 1, count[pair{key: "iii", data: 999}])
	})

	t.Run("for each with prefix not found", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("jjj", 101010)

		rmap.ForEachWithPrefix("k", func(_ string, _ interface{}) {
			panic("executing foreach with prefix not found")
		})
	})

	t.Run("for each with prefix found", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("kk0", 10100)
		rmap.Add("kk1", 10101)
		rmap.Add("kk2", 10102)

		count := map[pair]int{}
		rmap.ForEachWithPrefix("k", func(s string, data interface{}) {
			count[pair{key: s, data: data}]++
		})

		assert.Len(t, count, 3)
		assert.EqualValues(t, 1, count[pair{key: "kk0", data: 10100}])
		assert.EqualValues(t, 1, count[pair{key: "kk1", data: 10101}])
		assert.EqualValues(t, 1, count[pair{key: "kk2", data: 10102}])
	})
}

func TestMapThorough(t *testing.T) {
	t.Parallel()

	t.Run("does not contain prefix if not added", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("catastrophic", 0)

		assert.EqualValues(t, 1, rmap.Size())

		data, exists := rmap.Get("catastrophic")
		assert.True(t, exists)
		assert.EqualValues(t, 0, data)

		_, exists = rmap.Get("cat")
		assert.False(t, exists)
	})

	t.Run("contains prefix if added", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("catastrophic", 0)
		rmap.Add("cat", 1)

		assert.EqualValues(t, 2, rmap.Size())

		data, exists := rmap.Get("catastrophic")
		assert.True(t, exists)
		assert.EqualValues(t, 0, data)

		data, exists = rmap.Get("cat")
		assert.True(t, exists)
		assert.EqualValues(t, 1, data)
	})

	t.Run("does not contain common prefix if not added", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("butterfly", 22)
		rmap.Add("butterscotch", 33)

		assert.EqualValues(t, 2, rmap.Size())

		data, exists := rmap.Get("butterfly")
		assert.True(t, exists)
		assert.EqualValues(t, 22, data)

		data, exists = rmap.Get("butterscotch")
		assert.True(t, exists)
		assert.EqualValues(t, 33, data)

		_, exists = rmap.Get("butter")
		assert.False(t, exists)

		_, exists = rmap.Get("but")
		assert.False(t, exists)
	})

	t.Run("contains common prefix if added", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("butterfly", 22)
		rmap.Add("butterscotch", 33)
		rmap.Add("butter", 44)

		assert.EqualValues(t, 3, rmap.Size())

		data, exists := rmap.Get("butterfly")
		assert.True(t, exists)
		assert.EqualValues(t, 22, data)

		data, exists = rmap.Get("butterscotch")
		assert.True(t, exists)
		assert.EqualValues(t, 33, data)

		data, exists = rmap.Get("butter")
		assert.True(t, exists)
		assert.EqualValues(t, 44, data)
	})

	t.Run("contains word and prefix if added after prefix", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("trust", 5)
		rmap.Add("trustworthy", 6)

		assert.EqualValues(t, 2, rmap.Size())

		data, exists := rmap.Get("trust")
		assert.True(t, exists)
		assert.EqualValues(t, 5, data)

		data, exists = rmap.Get("trustworthy")
		assert.True(t, exists)
		assert.EqualValues(t, 6, data)
	})

	t.Run("removing prefix does not remove word", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("bittersweet", 100)
		rmap.Remove("bitter")

		assert.EqualValues(t, 1, rmap.Size())

		data, exists := rmap.Get("bittersweet")
		assert.True(t, exists)
		assert.EqualValues(t, 100, data)

		_, exists = rmap.Get("bitter")
		assert.False(t, exists)
	})

	t.Run("removing word does not remove prefix", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("bitter", 100)
		rmap.Remove("bittersweet")

		assert.EqualValues(t, 1, rmap.Size())

		data, exists := rmap.Get("bitter")
		assert.True(t, exists)
		assert.EqualValues(t, 100, data)

		_, exists = rmap.Get("bittersweet")
		assert.False(t, exists)
	})

	t.Run("remove common prefix does not remove words", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("hearing", 50)
		rmap.Add("heartless", 55)

		rmap.Remove("hear")

		assert.EqualValues(t, 2, rmap.Size())

		data, exists := rmap.Get("hearing")
		assert.True(t, exists)
		assert.EqualValues(t, 50, data)

		data, exists = rmap.Get("heartless")
		assert.True(t, exists)
		assert.EqualValues(t, 55, data)

		_, exists = rmap.Get("hear")
		assert.False(t, exists)
	})

	t.Run("remove existing common prefix", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("hearing", 123)
		rmap.Add("hear", 456)
		rmap.Add("heartless", 789)

		rmap.Remove("hear")

		assert.EqualValues(t, 2, rmap.Size())

		data, exists := rmap.Get("hearing")
		assert.True(t, exists)
		assert.EqualValues(t, 123, data)

		data, exists = rmap.Get("heartless")
		assert.True(t, exists)
		assert.EqualValues(t, 789, data)

		_, exists = rmap.Get("hear")
		assert.False(t, exists)
	})

	t.Run("add word with two existing prefixes", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("bliss", 1001)
		rmap.Add("blissful", 2002)
		rmap.Add("blissfulness", 3003)

		assert.EqualValues(t, 3, rmap.Size())

		data, exists := rmap.Get("bliss")
		assert.True(t, exists)
		assert.EqualValues(t, 1001, data)

		data, exists = rmap.Get("blissful")
		assert.True(t, exists)
		assert.EqualValues(t, 2002, data)

		data, exists = rmap.Get("blissfulness")
		assert.True(t, exists)
		assert.EqualValues(t, 3003, data)
	})

	t.Run("remove middle prefix of existing word", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("bliss", 81)
		rmap.Add("blissful", 90)
		rmap.Add("blissfulness", 99)

		rmap.Remove("blissful")

		assert.EqualValues(t, 2, rmap.Size())

		data, exists := rmap.Get("bliss")
		assert.True(t, exists)
		assert.EqualValues(t, 81, data)

		_, exists = rmap.Get("blissful")
		assert.False(t, exists)

		data, exists = rmap.Get("blissfulness")
		assert.True(t, exists)
		assert.EqualValues(t, 99, data)
	})

	t.Run("remove existing word with existing prefix", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("work", 35)
		rmap.Add("worker", 57)

		rmap.Remove("worker")

		assert.EqualValues(t, 1, rmap.Size())

		data, exists := rmap.Get("work")
		assert.True(t, exists)
		assert.EqualValues(t, 35, data)

		_, exists = rmap.Get("worker")
		assert.False(t, exists)
	})

	t.Run("remove existing word with existing common prefix", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("worker", 57)
		rmap.Add("workaholic", 79)

		rmap.Remove("worker")

		assert.EqualValues(t, 1, rmap.Size())

		_, exists := rmap.Get("work")
		assert.False(t, exists)

		_, exists = rmap.Get("worker")
		assert.False(t, exists)

		data, exists := rmap.Get("workaholic")
		assert.True(t, exists)
		assert.EqualValues(t, 79, data)
	})

	t.Run("for each with prefix when empty", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.ForEachWithPrefix("abcdef", func(_ string, _ interface{}) {
			panic("executing foreach with prefix on empty rmap")
		})
	})

	t.Run("for each with existing middle prefix", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("arm", 21)
		rmap.Add("armor", 23)
		rmap.Add("armored", 25)

		count := map[pair]int{}
		rmap.ForEachWithPrefix("armor", func(s string, data interface{}) {
			count[pair{key: s, data: data}]++
		})

		assert.EqualValues(t, 0, count[pair{key: "arm", data: 21}])
		assert.EqualValues(t, 1, count[pair{key: "armor", data: 23}])
		assert.EqualValues(t, 1, count[pair{key: "armored", data: 25}])
	})

	t.Run("for each with not existing prefix but existing prefix of prefix", func(t *testing.T) {
		rmap := radixtree.Map{}

		rmap.Add("arm", 21)
		rmap.Add("armored", 25)

		rmap.ForEachWithPrefix("armenia", func(_ string, _ interface{}) {
			panic("executing foreach with prefix not found")
		})
	})
}
