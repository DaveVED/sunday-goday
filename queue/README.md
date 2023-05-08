# queue
Queue data structure written in GO.

## What is a stack
Like stack, queue is a linear data structure that stores items in `First In First Out (FIFO)` manner. With a queue the least recently added item is removed first. A good example of queue is any queue of consumers for a resource where the consumer that came first is served first.

## Features
- [x] enqueue() - Adds an item to the queue. If the queue is full, then it is said to be an Overflow condition – Time Complexity `O(1)`
- [x] dequeue() - Removes an item from the queue. The items are popped in the same order in which they are pushed. If the queue is empty, then it is said to be an Underflow condition – Time Complexity `O(1)`
- [x] front() - Get the front item from queue – Time Complexity `O(1)`
- [x] rear() - Get the last item from queue – Time Complexity `O(1)`
- [x] size() - Get size of queue - Time Complexity `O(1)`
- [x] full() - Check if queue is full - Time Complexity `O(1)`
- [x] isEmpty() - check if queu is empty - Time Complexity `O(1)`