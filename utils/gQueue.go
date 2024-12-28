package utils

// This is a re-writing of the Queue function to allow for Go Generics,
// which allows me to add _any_ type of data to a queue, and allow the
// code to utilize it without the need for a wrapper function to
// enforce a datatype over the interfaces{} shortcut

// I re-wrote this because if I changed the older one it would break
// other challenges before I understood this concept

type GQueue[T any] struct {
	elements []T
}

// NewGQueue() creates a new instace of generic queue
func NewGQueue[T any]() GQueue[T] {
	return GQueue[T]{}
}

// Enqueue adds an element to the queue
func (q *GQueue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

// Dequeue removes and returns the first element from the queue
// returns a zero value of T if the queue is empty
func (q *GQueue[T]) Dequeue() (T, bool) {
	if len(q.elements) == 0 {
		var zero T
		return zero, false
	}

	bottom := q.elements[0]
	q.elements = q.elements[1:]

	return bottom, true
}

// Peek returns jthe first element without removing it
// returns the zero value of T if the queue is empty
func (q *GQueue[T]) Peek() (T, bool) {
	if len(q.elements) == 0 {
		var zero T
		return zero, false
	}

	return q.elements[0], true
}

// checks if queue is empty
func (q *GQueue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}
