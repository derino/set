package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyElement struct {
	id int
}

func MyHash(m *MyElement) int {
	return m.id
}

type YourElement struct {
	id string
}

func YourHash(y YourElement) string {
	return y.id
}

func TestSetWithPointerType(t *testing.T) {
	s := NewEmptySet(MyHash)

	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})
	assert.Equal(t, 3, s.Len())
	assert.True(t, s.Has(&MyElement{id: 1}))
	assert.True(t, s.Has(&MyElement{id: 2}))
	assert.True(t, s.Has(&MyElement{id: 3}))

	s.Remove(&MyElement{id: 2})
	assert.Equal(t, 2, s.Len())
	assert.True(t, s.Has(&MyElement{id: 1}))
	assert.False(t, s.Has(&MyElement{id: 2}))
	assert.True(t, s.Has(&MyElement{id: 3}))
}

func TestSetWithValueTypeAndStringHashType(t *testing.T) {
	s := NewEmptySet(YourHash)
	s.Add(YourElement{id: "1"})
	s.Add(YourElement{id: "2"})
	s.Add(YourElement{id: "2"})
	s.Add(YourElement{id: "3"})
	assert.Equal(t, 3, s.Len())
}

func TestTypeAliasAndToList(t *testing.T) {
	type MyElementSet = Set[int, *MyElement]

	myElements := []*MyElement{{id: 1}, {id: 2}}
	s := MyElementSet{Elems: make(map[int]*MyElement), hasher: MyHash}
	for _, e := range myElements {
		s.Add(e)
	}

	assert.ElementsMatch(t, myElements, s.ToList())
}

func TestNewSet(t *testing.T) {
	elems := []*MyElement{{id: 1}, {id: 2}, {id: 2}, {id: 3}}
	s := NewSet(elems, MyHash)
	assert.Equal(t, 3, s.Len())
	assert.True(t, s.Has(&MyElement{id: 1}))
	assert.True(t, s.Has(&MyElement{id: 2}))
	assert.True(t, s.Has(&MyElement{id: 3}))
}

func TestClone(t *testing.T) {
	s := NewEmptySet(MyHash)
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := s.Clone()
	assert.Equal(t, 3, s2.Len())
	assert.True(t, s2.Has(&MyElement{id: 1}))
	assert.True(t, s2.Has(&MyElement{id: 2}))
	assert.True(t, s2.Has(&MyElement{id: 3}))
}

func TestUnion(t *testing.T) {
	s := NewEmptySet(MyHash)
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := NewEmptySet(MyHash)
	s2.Add(&MyElement{id: 4})
	s2.Add(&MyElement{id: 3})

	union := Union(s, s2)
	assert.Equal(t, 4, union.Len())
	assert.True(t, union.Has(&MyElement{id: 1}))
	assert.True(t, union.Has(&MyElement{id: 2}))
	assert.True(t, union.Has(&MyElement{id: 3}))
	assert.True(t, union.Has(&MyElement{id: 4}))

	s.Union(s2)
	assert.Equal(t, 4, s.Len())
	assert.True(t, s.Has(&MyElement{id: 1}))
	assert.True(t, s.Has(&MyElement{id: 2}))
	assert.True(t, s.Has(&MyElement{id: 3}))
	assert.True(t, s.Has(&MyElement{id: 4}))
}

func TestIntersect(t *testing.T) {
	s := NewEmptySet(MyHash)
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := NewEmptySet(MyHash)
	s2.Add(&MyElement{id: 4})
	s2.Add(&MyElement{id: 3})

	isect := Intersect(s, s2)
	assert.Equal(t, 1, isect.Len())
	assert.True(t, isect.Has(&MyElement{id: 3}))

	s.Intersect(s2)
	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Has(&MyElement{id: 3}))
}

func TestDifference(t *testing.T) {
	s := NewEmptySet(MyHash)
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := NewEmptySet(MyHash)
	s2.Add(&MyElement{id: 4})
	s2.Add(&MyElement{id: 3})

	diff := Difference(s, s2)
	assert.Equal(t, 2, diff.Len())
	assert.True(t, diff.Has(&MyElement{id: 1}))
	assert.True(t, diff.Has(&MyElement{id: 2}))

	s.Difference(s2)
	assert.Equal(t, 2, s.Len())
	assert.True(t, s.Has(&MyElement{id: 1}))
	assert.True(t, s.Has(&MyElement{id: 2}))
}

func TestIsSubsetAndIsSuperset(t *testing.T) {
	s := NewEmptySet(MyHash)
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := NewEmptySet(MyHash)
	s2.Add(&MyElement{id: 4})
	s2.Add(&MyElement{id: 3})

	s3 := NewEmptySet(MyHash)
	s3.Add(&MyElement{id: 2})
	s3.Add(&MyElement{id: 3})

	empty := NewEmptySet(MyHash)

	assert.True(t, empty.IsSubset(s))
	assert.False(t, s.IsSubset(empty))
	assert.False(t, empty.IsSuperset(s))
	assert.True(t, s.IsSuperset(empty))

	assert.True(t, s.IsSubset(s))
	assert.True(t, s.IsSuperset(s))

	assert.False(t, s.IsSubset(s2))
	assert.False(t, s2.IsSubset(s))
	assert.False(t, s.IsSuperset(s2))
	assert.False(t, s2.IsSuperset(s))

	assert.True(t, s3.IsSubset(s))
	assert.False(t, s.IsSubset(s3))
	assert.True(t, s.IsSuperset(s3))
	assert.False(t, s3.IsSuperset(s))

	assert.True(t, IsSubset(s3, s))
	assert.False(t, IsSubset(s, s3))
	assert.True(t, IsSuperset(s, s3))
	assert.False(t, IsSuperset(s3, s))
}

func TestIsDisjoint(t *testing.T) {
	s := NewEmptySet(MyHash)
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := NewEmptySet(MyHash)
	s2.Add(&MyElement{id: 4})
	s2.Add(&MyElement{id: 3})

	s3 := NewEmptySet(MyHash)
	s3.Add(&MyElement{id: 5})
	s3.Add(&MyElement{id: 6})

	empty := NewEmptySet(MyHash)

	assert.True(t, empty.IsDisjoint(s))
	assert.True(t, s.IsDisjoint(empty))
	assert.False(t, s.IsDisjoint(s))

	assert.False(t, s.IsDisjoint(s2))
	assert.False(t, s2.IsDisjoint(s))
	assert.True(t, s3.IsDisjoint(s))
	assert.True(t, s.IsDisjoint(s3))

	assert.False(t, IsDisjoint(s, s2))
	assert.True(t, IsDisjoint(s, s3))
}

func TestEqual(t *testing.T) {
	s := NewEmptySet(MyHash)
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := NewEmptySet(MyHash)
	s2.Add(&MyElement{id: 2})
	s2.Add(&MyElement{id: 3})

	s3 := NewEmptySet(MyHash)
	s3.Add(&MyElement{id: 1})
	s3.Add(&MyElement{id: 2})
	s3.Add(&MyElement{id: 3})

	s4 := NewEmptySet(MyHash)
	s4.Add(&MyElement{id: 1})
	s4.Add(&MyElement{id: 2})
	s4.Add(&MyElement{id: 4})

	empty := NewEmptySet(MyHash)

	assert.False(t, empty.Equal(s))
	assert.False(t, s.Equal(empty))

	assert.False(t, s.Equal(s2))
	assert.False(t, s2.Equal(s))

	assert.False(t, s.Equal(s4))
	assert.False(t, s4.Equal(s))

	assert.True(t, s.Equal(s3))
	assert.True(t, s3.Equal(s))

	assert.True(t, Equal(s, s3))
	assert.False(t, Equal(s, s2))
}
