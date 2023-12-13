package heapmap

import (
	"testing"
)

type person struct {
	id      uint
	name    string
	address string
	age     int
}

var (
	john = person{1, "John", "123 Main St", 42}
	jane = person{2, "Jane", "456 Main St", 43}
	joe  = person{3, "Joe", "789 Main St", 44}
)

func TestContents(t *testing.T) {
	testLen := func(hm HeapMap[uint, person, int], expected int, details string) {
		t.Run("Len", func(t *testing.T) {
			if hm.Len() != expected {
				t.Errorf("Len() = %d, want %d (%s)", hm.Len(), expected, details)
			}
		})
	}

	testEmpty := func(hm HeapMap[uint, person, int], expected bool, details string) {
		t.Run("Empty", func(t *testing.T) {
			if hm.Empty() != expected {
				t.Errorf("Empty() = %t, want %t (%s)", hm.Empty(), expected, details)
			}
		})
	}

	testPeek := func(hm HeapMap[uint, person, int], expected person, expectedOk bool, details string) {
		t.Run("Peek", func(t *testing.T) {
			e, ok := hm.Peek()
			if ok != expectedOk {
				t.Errorf("Peek() = %v, %t, want %v, %t (%s)", e, ok, expected, expectedOk, details)
			}
			if ok && e.Value != expected {
				t.Errorf("Peek() = %v, %t, want %v, %t (%s)", e, ok, expected, expectedOk, details)
			}
		})
	}

	testPop := func(hm HeapMap[uint, person, int], expected person, expectedOk bool, details string) {
		t.Run("Pop", func(t *testing.T) {
			e, ok := hm.Pop()
			if ok != expectedOk {
				t.Errorf("Pop() = %v, %t, want %v, %t (%s)", e, ok, expected, expectedOk, details)
			}
			if ok && e.Value != expected {
				t.Errorf("Pop() = %v, %t, want %v, %t (%s)", e, ok, expected, expectedOk, details)
			}
		})
	}

	testContains := func(hm HeapMap[uint, person, int], key uint, expectedOk bool, details string) {
		t.Run("Contains", func(t *testing.T) {
			ok := hm.Contains(key)
			if ok != expectedOk {
				t.Errorf("Contains(%d) = %t, want %t (%s)", key, ok, expectedOk, details)
			}
		})
	}

	testGet := func(hm HeapMap[uint, person, int], key uint, expected person, expectedOk bool, details string) {
		t.Run("Get", func(t *testing.T) {
			e, ok := hm.Get(key)
			if ok != expectedOk {
				t.Errorf("Get(%d) = %v, %t, want %v, %t (%s)", key, e, ok, expected, expectedOk, details)

			}
			if ok && e.Value != expected {
				t.Errorf("Get(%d) = %v, %t, want %v, %t (%s)", key, e, ok, expected, expectedOk, details)
			}
		})
	}

	hm := NewMin[uint, person, int]()
	testLen(hm, 0, "after NewMin")
	testEmpty(hm, true, "after NewMin")
	testContains(hm, john.id, false, "after NewMin")
	testGet(hm, john.id, person{}, false, "after NewMin")
	testPeek(hm, person{}, false, "after NewMin")
	testPop(hm, person{}, false, "after NewMin")

	hm.Set(john.id, john, john.age)
	hm.Set(jane.id, jane, jane.age)
	hm.Set(joe.id, joe, joe.age)
	testLen(hm, 3, "after Set")
	testEmpty(hm, false, "after Set")
	testContains(hm, john.id, true, "after Set")
	testGet(hm, john.id, john, true, "after Set")
	testPeek(hm, john, true, "after Set")
	testPop(hm, john, true, "after Set")
	testContains(hm, john.id, false, "after Set")
	testGet(hm, john.id, person{}, false, "after Set")

	testContains(hm, jane.id, true, "before Remove")
	testGet(hm, jane.id, jane, true, "before Remove")
	hm.Remove(jane.id)
	testContains(hm, jane.id, false, "after Remove")
	testGet(hm, jane.id, person{}, false, "after Remove")

	testContains(hm, joe.id, true, "before Clear")
	testGet(hm, joe.id, joe, true, "before Clear")
	hm.Clear()
	testContains(hm, joe.id, false, "after Clear")
	testGet(hm, joe.id, person{}, false, "after Clear")

	testLen(hm, 0, "after Clear")
	testEmpty(hm, true, "after Clear")
}

func TestMinVsMax(t *testing.T) {
	hmMin := NewMin[uint, person, int]()
	hmMin.Set(john.id, john, john.age)
	hmMin.Set(jane.id, jane, jane.age)
	hmMin.Set(joe.id, joe, joe.age)

	hmMax := NewMax[uint, person, int]()
	hmMax.Set(john.id, john, john.age)
	hmMax.Set(jane.id, jane, jane.age)
	hmMax.Set(joe.id, joe, joe.age)

	if e, ok := hmMin.Peek(); !ok || e.Value != john {
		t.Errorf("hmMin.Peek() = %v, %t, want %v, %t", e, ok, john, true)
	}

	if e, ok := hmMax.Peek(); !ok || e.Value != joe {
		t.Errorf("hmMax.Peek() = %v, %t, want %v, %t", e, ok, joe, true)
	}
}

func TestSetOverride(t *testing.T) {
	hm := NewMin[uint, person, int]()

	hm.Set(john.id, john, john.age)
	if e, ok := hm.Get(john.id); !ok || e.Value != john {
		t.Errorf("hm.Get(%d) = %v, %t, want %v, %t", john.id, e, ok, john, true)
	}

	hm.Set(john.id, jane, jane.age)
	if e, ok := hm.Get(john.id); !ok || e.Value != jane {
		t.Errorf("hm.Get(%d) = %v, %t, want %v, %t", john.id, e, ok, jane, true)
	}
}
