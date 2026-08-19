package graph

import (
	"errors"
)

// Queue represents a First-In-First-Out data structure using generics.
type Queue[T any] struct {
	elements []T
}

// Enqueue adds an element to the back of the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the front element of the queue.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	element := q.elements[0]

	// Prevent memory leaks by zeroing out the reference before slicing
	var zero T
	q.elements[0] = zero

	q.elements = q.elements[1:]
	return element, nil
}

// IsEmpty checks if the queue has no items.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) Elements() []T {
	dst := make([]T, len(q.elements))
	copy(dst, q.elements)
	return dst
}
