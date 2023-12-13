package heapmap

import "fmt"

func Example() {
	// Create a new min-heap map.
	hmMin := NewMin[uint, string, int]()

	// Insert some entries.
	hmMin.Set(1, "John", 42)
	hmMin.Set(2, "Jane", 43)
	hmMin.Set(3, "Joe", 44)

	// Peek at the entry with the highest priority.
	e, ok := hmMin.Peek()
	if ok {
		fmt.Printf("%d: %s\n", e.Key, e.Value)
	}

	// Remove the entry with the highest priority.
	e, ok = hmMin.Pop()
	if ok {
		fmt.Printf("%d: %s\n", e.Key, e.Value)
	}

	// Check the length of the heap map.
	fmt.Printf("Len: %d\n", hmMin.Len())

	// Check if the heap map is empty.
	fmt.Printf("Empty: %t\n", hmMin.Empty())

	// Check if the heap map contains a key.
	fmt.Printf("Contains(1): %t\n", hmMin.Contains(1))
	fmt.Printf("Contains(2): %t\n", hmMin.Contains(2))
	fmt.Printf("Contains(3): %t\n", hmMin.Contains(3))

	// Overwrite an entry.
	hmMin.Set(3, "Joey", 45)

	// Get the new entry.
	e, ok = hmMin.Get(3)
	if ok {
		fmt.Printf("%d: %s\n", e.Key, e.Value)
	}

	// Create a new max-heap map.
	hmMax := NewMax[uint, string, int]()

	// Insert all entries from the min-heap map into the max-heap map.
	for _, e := range hmMin.Entries() {
		hmMax.Set(e.Key, e.Value, e.Priority)
	}

	// Get the entry with the highest priority.
	e, ok = hmMax.Peek()
	if ok {
		fmt.Printf("%d: %s\n", e.Key, e.Value)
	}

	// Remove the entry with key 3.
	hmMax.Remove(3)

	// Get the entry with the highest priority.
	e, ok = hmMax.Peek()
	if ok {
		fmt.Printf("%d: %s\n", e.Key, e.Value)
	}

	// Clear the heap map.
	hmMax.Clear()

	// Check the length of the heap map.
	fmt.Printf("Len: %d\n", hmMax.Len())

	// Output:
	// 1: John
	// 1: John
	// Len: 2
	// Empty: false
	// Contains(1): false
	// Contains(2): true
	// Contains(3): true
	// 3: Joey
	// 3: Joey
	// 2: Jane
	// Len: 0
}
