package heapmap

import (
	"cmp"
	"container/heap"
)

type Entry[K comparable, V, P any] struct {
	Key      K
	Value    V
	Priority P
	index    int
}

type HeapMap[K comparable, V, P any] interface {
	Len() int
	Empty() bool

	Peek() (Entry[K, V, P], bool)
	Pop() (Entry[K, V, P], bool)

	Get(key K) (Entry[K, V, P], bool)
	Set(key K, value V, priority P)

	Contains(key K) bool
	Remove(key K)
	Clear()

	Keys() []K
	Values() []V
	Entries() []Entry[K, V, P]
}

func New[K comparable, V, P any](less func(P, P) bool) HeapMap[K, V, P] {
	return &heapmap[K, V, P]{
		h: pq[K, V, P]{
			less: less,
		},
		m: make(map[K]*Entry[K, V, P]),
	}
}

func NewMin[K comparable, V any, P cmp.Ordered]() HeapMap[K, V, P] {
	return New[K, V, P](func(a, b P) bool {
		return a < b
	})
}

func NewMax[K comparable, V any, P cmp.Ordered]() HeapMap[K, V, P] {
	return New[K, V, P](func(a, b P) bool {
		return a > b
	})
}

type heapmap[K comparable, V, P any] struct {
	h pq[K, V, P]
	m map[K]*Entry[K, V, P]
}

func (hm *heapmap[K, V, P]) Len() int {
	return len(hm.m)
}

func (hm *heapmap[K, V, P]) Empty() bool {
	return hm.Len() == 0
}

func (hm *heapmap[K, V, P]) Peek() (Entry[K, V, P], bool) {
	if hm.Empty() {
		return Entry[K, V, P]{}, false
	}
	return *hm.h.entries[0], true
}

func (hm *heapmap[K, V, P]) Pop() (Entry[K, V, P], bool) {
	if hm.Empty() {
		return Entry[K, V, P]{}, false
	}
	e := *heap.Pop(&hm.h).(*Entry[K, V, P])
	delete(hm.m, e.Key)
	return e, true
}

func (hm *heapmap[K, V, P]) Get(key K) (Entry[K, V, P], bool) {
	if e, ok := hm.m[key]; ok {
		return *e, true
	}
	return Entry[K, V, P]{}, false
}

func (hm *heapmap[K, V, P]) Set(key K, value V, priority P) {
	if e, ok := hm.m[key]; ok {
		e.Value = value
		e.Priority = priority
		heap.Fix(&hm.h, e.index)
		return
	}
	e := &Entry[K, V, P]{
		Key:      key,
		Value:    value,
		Priority: priority,
	}
	heap.Push(&hm.h, e)
	hm.m[key] = e
}

func (hm *heapmap[K, V, P]) Contains(key K) bool {
	_, ok := hm.m[key]
	return ok
}

func (hm *heapmap[K, V, P]) Remove(key K) {
	e, ok := hm.m[key]
	if !ok {
		return
	}
	heap.Remove(&hm.h, e.index)
	delete(hm.m, key)
}

func (hm *heapmap[K, V, P]) Clear() {
	hm.h.entries = hm.h.entries[:0]
	hm.m = make(map[K]*Entry[K, V, P])
}

func (hm *heapmap[K, V, P]) Keys() []K {
	keys := make([]K, 0, hm.Len())
	for key := range hm.m {
		keys = append(keys, key)
	}
	return keys
}

func (hm *heapmap[K, V, P]) Values() []V {
	values := make([]V, 0, hm.Len())
	for _, e := range hm.m {
		values = append(values, e.Value)
	}
	return values
}

func (hm *heapmap[K, V, P]) Entries() []Entry[K, V, P] {
	entries := make([]Entry[K, V, P], 0, hm.Len())
	for _, e := range hm.m {
		entries = append(entries, *e)
	}
	return entries
}
