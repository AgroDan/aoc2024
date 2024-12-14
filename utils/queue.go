package utils

// a LIFO queue

type Queue struct {
	elements []interface{}
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Enqueue(element interface{}) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.elements) == 0 {
		return nil
	}
	bottom := q.elements[0]
	q.elements = q.elements[1:]
	return bottom
}

func (q *Queue) Peek() interface{} {
	if len(q.elements) == 0 {
		return nil
	}
	return q.elements[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}
