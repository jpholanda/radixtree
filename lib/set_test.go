package radixtree_test

import (
	"testing"

	radixtree "github.com/jpholanda/radixtree/lib"
	"github.com/stretchr/testify/assert"
)

func TestSetBasic(t *testing.T) {
	t.Parallel()

	t.Run("empty contains nothing", func(t *testing.T) {
		set := radixtree.Set{}

		assert.False(t, set.Contains("aaa"))
	})

	t.Run("contain after add", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("bbb")

		assert.True(t, set.Contains("bbb"))
	})

	t.Run("not contain after add and remove", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("ccc")
		set.Remove("ccc")

		assert.False(t, set.Contains("ccc"))
	})

	t.Run("size after add", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("ddd")

		assert.EqualValues(t, 1, set.Size())
	})

	t.Run("size after remove", func(t *testing.T) {
		set := radixtree.Set{}

		set.Remove("eee")

		assert.EqualValues(t, 0, set.Size())
	})

	t.Run("size after add then remove", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("fff")
		set.Remove("fff")

		assert.EqualValues(t, 0, set.Size())
	})

	t.Run("for each when empty", func(t *testing.T) {
		set := radixtree.Set{}

		set.ForEach(func(_ string) {
			panic("should not be executed")
		})
	})

	t.Run("for each when not empty", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("ggg")
		set.Add("hhh")
		set.Add("iii")

		count := map[string]int{}
		set.ForEach(func(s string) {
			count[s]++
		})

		assert.Len(t, count, 3)
		assert.EqualValues(t, 1, count["ggg"])
		assert.EqualValues(t, 1, count["hhh"])
		assert.EqualValues(t, 1, count["iii"])
	})

	t.Run("for each with prefix not found", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("jjj")

		set.ForEachWithPrefix("k", func(_ string) {
			panic("executing foreach with prefix not found")
		})
	})

	t.Run("for each with prefix found", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("kk0")
		set.Add("kk1")
		set.Add("kk2")

		count := map[string]int{}
		set.ForEachWithPrefix("k", func(s string) {
			count[s]++
		})

		assert.Len(t, count, 3)
		assert.EqualValues(t, 1, count["kk0"])
		assert.EqualValues(t, 1, count["kk1"])
		assert.EqualValues(t, 1, count["kk2"])
	})
}

func TestSetThorough(t *testing.T) {
	t.Parallel()

	t.Run("does not contain prefix if not added", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("catastrophic")

		assert.EqualValues(t, 1, set.Size())
		assert.True(t, set.Contains("catastrophic"))
		assert.False(t, set.Contains("cat"))
	})

	t.Run("contains prefix if added", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("catastrophic")
		set.Add("cat")

		assert.EqualValues(t, 2, set.Size())
		assert.True(t, set.Contains("catastrophic"))
		assert.True(t, set.Contains("cat"))
	})

	t.Run("does not contain common prefix if not added", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("butterfly")
		set.Add("butterscotch")

		assert.EqualValues(t, 2, set.Size())
		assert.True(t, set.Contains("butterfly"))
		assert.True(t, set.Contains("butterscotch"))
		assert.False(t, set.Contains("butter"))
		assert.False(t, set.Contains("but"))
	})

	t.Run("contains common prefix if added", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("butterfly")
		set.Add("butterscotch")
		set.Add("butter")

		assert.EqualValues(t, 3, set.Size())
		assert.True(t, set.Contains("butterfly"))
		assert.True(t, set.Contains("butterscotch"))
		assert.True(t, set.Contains("butter"))
	})

	t.Run("contains word and prefix if added after prefix", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("trust")
		set.Add("trustworthy")

		assert.EqualValues(t, 2, set.Size())
		assert.True(t, set.Contains("trust"))
		assert.True(t, set.Contains("trustworthy"))
	})

	t.Run("removing prefix does not remove word", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("bittersweet")
		set.Remove("bitter")

		assert.EqualValues(t, 1, set.Size())
		assert.True(t, set.Contains("bittersweet"))
		assert.False(t, set.Contains("bitter"))
	})

	t.Run("removing word does not remove prefix", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("bitter")
		set.Remove("bittersweet")

		assert.EqualValues(t, 1, set.Size())
		assert.True(t, set.Contains("bitter"))
		assert.False(t, set.Contains("bittersweet"))
	})

	t.Run("remove common prefix does not remove words", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("hearing")
		set.Add("heartless")

		set.Remove("hear")

		assert.EqualValues(t, 2, set.Size())
		assert.True(t, set.Contains("hearing"))
		assert.True(t, set.Contains("heartless"))
		assert.False(t, set.Contains("hear"))
	})

	t.Run("remove existing common prefix", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("hearing")
		set.Add("hear")
		set.Add("heartless")

		set.Remove("hear")

		assert.EqualValues(t, 2, set.Size())
		assert.True(t, set.Contains("hearing"))
		assert.True(t, set.Contains("heartless"))
		assert.False(t, set.Contains("hear"))
	})

	t.Run("add word with two existing prefixes", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("bliss")
		set.Add("blissful")
		set.Add("blissfulness")

		assert.EqualValues(t, 3, set.Size())
		assert.True(t, set.Contains("bliss"))
		assert.True(t, set.Contains("blissful"))
		assert.True(t, set.Contains("blissfulness"))
	})

	t.Run("remove middle prefix of existing word", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("bliss")
		set.Add("blissful")
		set.Add("blissfulness")

		set.Remove("blissful")

		assert.EqualValues(t, 2, set.Size())
		assert.True(t, set.Contains("bliss"))
		assert.False(t, set.Contains("blissful"))
		assert.True(t, set.Contains("blissfulness"))
	})

	t.Run("remove existing word with existing prefix", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("work")
		set.Add("worker")

		set.Remove("worker")

		assert.EqualValues(t, 1, set.Size())
		assert.True(t, set.Contains("work"))
		assert.False(t, set.Contains("worker"))
	})

	t.Run("remove existing word with existing common prefix", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("worker")
		set.Add("workaholic")

		set.Remove("worker")

		assert.EqualValues(t, 1, set.Size())
		assert.False(t, set.Contains("work"))
		assert.False(t, set.Contains("worker"))
		assert.True(t, set.Contains("workaholic"))
	})

	t.Run("for each with prefix when empty", func(t *testing.T) {
		set := radixtree.Set{}

		set.ForEachWithPrefix("abcdef", func(_ string) {
			panic("executing foreach with prefix on empty set")
		})
	})

	t.Run("for each with existing middle prefix", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("arm")
		set.Add("armor")
		set.Add("armored")

		count := map[string]int{}
		set.ForEachWithPrefix("armor", func(s string) {
			count[s]++
		})

		assert.EqualValues(t, 1, count["armor"])
		assert.EqualValues(t, 1, count["armored"])
		assert.EqualValues(t, 0, count["armenia"])
	})

	t.Run("for each with not existing prefix but existing prefix of prefix", func(t *testing.T) {
		set := radixtree.Set{}

		set.Add("arm")
		set.Add("armored")

		set.ForEachWithPrefix("armenia", func(s string) {
			panic("executing foreach with prefix not found")
		})
	})
}
