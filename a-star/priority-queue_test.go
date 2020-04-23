package a_star

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	items := []Item{
		{Value: []int{1, 2}, Priority: 5},
		{Value: []int{3, 4}, Priority: 3},
		{Value: []int{5, 4}, Priority: 6},
		{Value: []int{1, 2}, Priority: 1},
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for _, item := range items {
		i := item
		heap.Push(&pq, &i)
	}

	heap.Push(&pq, &Item{Value: []int{4, 5}, Priority: 4})
	heap.Push(&pq, &Item{Value: []int{1, 23}, Priority: 2})

	//expectValueResult := []int{6, 5, 4, 3, 2, 1}
	expectValueResult := []int{1, 2, 3, 4, 5, 6}
	result := []int{}
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		result = append(result, item.Priority)
	}

	if !reflect.DeepEqual(result, expectValueResult) {
		t.Errorf("expected %v,\nbut got %v", expectValueResult, result)
	}
}
