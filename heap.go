package heapmap

type pq[K comparable, V, P any] struct {
	entries []*Entry[K, V, P]
	less    func(P, P) bool
}

func (h pq[K, V, P]) Len() int {
	return len(h.entries)
}

func (h pq[K, V, P]) Less(i, j int) bool {
	return h.less(h.entries[i].Priority, h.entries[j].Priority)
}

func (h pq[K, V, P]) Swap(i, j int) {
	h.entries[i], h.entries[j] = h.entries[j], h.entries[i]
	h.entries[i].index = i
	h.entries[j].index = j
}

func (h *pq[K, V, P]) Push(x any) {
	n := len(h.entries)
	e := x.(*Entry[K, V, P])
	e.index = n
	h.entries = append(h.entries, e)
}

func (h *pq[K, V, P]) Pop() any {
	n := len(h.entries)
	e := h.entries[n-1]
	h.entries = h.entries[0 : n-1]
	return e
}
