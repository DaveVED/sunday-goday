package main

import (
	"fmt"
)

func isEmpty(stack []int) bool {
	return len(stack) == 0
}

func size(stack []int) int {
	return len(stack)
}

func top(stack []int) int {
	if len(stack) <= 0 {
		return -1
	}

	return stack[0]
}

func push(a int, stack *[]int) {
	//create new slice, and append with old one.
	*stack = append([]int{a}, *stack...)
}

func pop(stack *[]int) {
	// uses a slice operator to trim first item.
	*stack = (*stack)[1:len(*stack)]
}

func dump(stack *[]int, a ...int) {
	/*
	* Wanted to play around with Variadic Functions - Function to be called with any number of trailing arguments.
	* To test, created this `dump` which is not normally in a stack, but for this case, will push a list of items to stack.
	 */
	temp := []int{}
	for _, v := range a {
		temp = append(temp, v)
	}

	*stack = append(temp, *stack...)
}

func main() {
	// init.
	stack := []int{}
	fmt.Printf("Init stack set to --> %v\n\n", stack)

	fmt.Println("Testing function base cases (empty stack)")
	fmt.Printf("  * Stack is empty --> %v\n", isEmpty(stack))
	fmt.Printf("  * Stack size is --> %v\n", size(stack))
	fmt.Printf("  * Stack top item is --> %v\n\n", top(stack))

	// easy stack setup.
	fmt.Printf("Setting up a basic stack with initial values --> 35, 12, & 74\n")
	push(35, &stack)
	push(12, &stack)
	push(74, &stack)
	fmt.Printf("  * Stack is --> %v\n\n", stack)

	fmt.Println("Poping item --> 74 (last in.)")
	pop(&stack)
	fmt.Printf("  * Stack is --> %v\n", stack)
	fmt.Printf("  * Stack is empty --> %v\n", isEmpty(stack))
	fmt.Printf("  * Stack size is --> %v\n", size(stack))
	fmt.Printf("  * Stack top item is --> %v\n\n", top(stack))

	fmt.Println("Adding in 6 random numbers.")
	push(17, &stack)
	push(53, &stack)
	push(88, &stack)
	push(91, &stack)
	push(44, &stack)
	push(2, &stack)
	fmt.Printf("  * Stack is --> %v\n\n", stack)

	fmt.Println("Removing 2 items.")
	pop(&stack)
	pop(&stack)
	fmt.Printf("  * Stack is --> %v\n", stack)
	fmt.Printf("  * Stack is empty --> %v\n", isEmpty(stack))
	fmt.Printf("  * Stack size is --> %v\n", size(stack))
	fmt.Printf("  * Stack top item is --> %v\n\n", top(stack))

	fmt.Println("Dumping 3 values.")
	dump(&stack, 99, 64, 3)
	fmt.Printf("  * Stack is --> %v\n", stack)
}
