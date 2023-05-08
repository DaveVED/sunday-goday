package main

import (
	"flag"
	"fmt"
)

func enqueue(value int, maxSize int, queue *[]int) {
	if !full(maxSize, *queue) {
		*queue = append([]int{value}, *queue...)
	} else {
		fmt.Println("Queue is full, can't add.")
	}

}

func dequeue(queue *[]int) {
	*queue = (*queue)[:len(*queue)-1]
}

func front(queue []int) int {
	if len(queue) <= 0 {
		return -1
	}

	return queue[0]
}

func rear(queue []int) int {
	if len(queue) <= 0 {
		return -1
	}

	return queue[len(queue)-1]
}

func size(queue []int) int {
	return len(queue)
}

func full(maxSize int, queue []int) bool {
	return len(queue) >= maxSize
}

func isEmpty(queue []int) bool {
	return len(queue) == 0
}

func main() {
	maxSize := flag.Int("size", 10, "Queue max size.")
	flag.Parse()

	queue := []int{}
	fmt.Printf("Init queue set to --> %v\n\n", queue)

	fmt.Println("Testing function base cases (empty queue)")
	fmt.Printf("  * Queue is empty --> %v\n", isEmpty(queue))
	fmt.Printf("  * Queue size is --> %v\n", size(queue))
	fmt.Printf("  * Queue size full --> %v\n", full(*maxSize, queue))
	fmt.Printf("  * Queue rear item is --> %v\n", rear(queue))
	fmt.Printf("  * Queue front item is --> %v\n\n", front(queue))

	fmt.Printf("Setting up a basic queue with initial values --> 3, 17, 24\n")
	enqueue(3, *maxSize, &queue)
	enqueue(17, *maxSize, &queue)
	enqueue(24, *maxSize, &queue)
	fmt.Printf("  * Queue is --> %v\n\n", queue)
	fmt.Printf("  * Queue is empty --> %v\n", isEmpty(queue))
	fmt.Printf("  * Queue size is --> %v\n", size(queue))
	fmt.Printf("  * Queue size full --> %v\n", full(*maxSize, queue))
	fmt.Printf("  * Queue rear item is --> %v\n", rear(queue))
	fmt.Printf("  * Queue front item is --> %v\n\n", front(queue))

	dequeue(&queue)
	fmt.Println("Queue item --> 74 (frist in.)")
	fmt.Printf("  * Queue is --> %v\n", queue)
	fmt.Printf("  * Queue is empty --> %v\n", isEmpty(queue))
	fmt.Printf("  * Queue size is --> %v\n", size(queue))
	fmt.Printf("  * Queue size full --> %v\n", full(*maxSize, queue))
	fmt.Printf("  * Queue rear item is --> %v\n", rear(queue))
	fmt.Printf("  * Queue front item is --> %v\n\n", front(queue))

	fmt.Printf("Adding in enough to create a 'full' queue (6 more itmes for my test)\n")
	enqueue(44, *maxSize, &queue)
	enqueue(68, *maxSize, &queue)
	enqueue(99, *maxSize, &queue)
	enqueue(81, *maxSize, &queue)
	enqueue(25, *maxSize, &queue)
	enqueue(77, *maxSize, &queue)
	enqueue(12, *maxSize, &queue)
	enqueue(60, *maxSize, &queue)
	enqueue(89, *maxSize, &queue)
	fmt.Printf("  * Queue is --> %v\n\n", queue)
	fmt.Printf("  * Queue is empty --> %v\n", isEmpty(queue))
	fmt.Printf("  * Queue size is --> %v\n", size(queue))
	fmt.Printf("  * Queue size full --> %v\n", full(*maxSize, queue))
	fmt.Printf("  * Queue rear item is --> %v\n", rear(queue))
	fmt.Printf("  * Queue front item is --> %v\n\n", front(queue))

	fmt.Println("Removing two items, to ensure it's not full anymore")
	dequeue(&queue)
	dequeue(&queue)
	fmt.Printf("  * Queue is --> %v\n\n", queue)
	fmt.Printf("  * Queue is empty --> %v\n", isEmpty(queue))
	fmt.Printf("  * Queue size is --> %v\n", size(queue))
	fmt.Printf("  * Queue size full --> %v\n", full(*maxSize, queue))
	fmt.Printf("  * Queue rear item is --> %v\n", rear(queue))
	fmt.Printf("  * Queue front item is --> %v\n\n", front(queue))
}
