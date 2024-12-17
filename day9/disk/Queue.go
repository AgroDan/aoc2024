package disk

import "errors"

// Not sure how best to do that so I'm just going to make a
// queue specific to the File object.

type Queue struct {
	elements []File
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Enqueue(element File) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() (File, error) {
	if len(q.elements) == 0 {
		return File{}, errors.New("empty queue")
	}
	bottom := q.elements[0]
	q.elements = q.elements[1:]
	return bottom, nil
}

func (q *Queue) Peek() (File, error) {
	if len(q.elements) == 0 {
		return File{}, errors.New("empty queue")
	}
	return q.elements[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}
