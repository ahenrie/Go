package main

import "fmt"

//represent class for queue
type Queue struct {
	nodes []interface{}
}

//add something 
func (q *Queue) add(node interface{}) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) isEmpty() bool {
	return len(q.nodes) == 0
}

func (q *Queue) get() interface{} {
	if len(q.nodes) == 0 {
		return nil
	}

	//return first node
	return q.nodes[0]
}

func (q *Queue) remove() interface{} {
	if len(q.nodes) == 0 {
		return nil
	}

	node := q.nodes[0]
	q.nodes = q.nodes[1:]

	return node
}

func (q *Queue) size() interface{} {
	return len(q.nodes)
}

func (q *Queue) clear() {
	q.nodes = nil
}

func (q *Queue) changeImpl(newImplementation *Queue) {
	tempQueue := Queue{}

	//transfer to temp
	for !q.isEmpty() {
		newNode := q.remove()
		tempQueue.add(newNode)
	}

	//replace
	*q = *newImplementation

	//transfer to new
	for !tempQueue.isEmpty() {
		newNode := tempQueue.remove()
		q.add(newNode)
	}
}

func main() {
	// Ints
	fmt.Println("*****************INT*****************")
	intQueue := Queue{}

	intQueue.add(100)
	intQueue.add(2)
	intQueue.add(22)

	fmt.Println("Original Queue:")
	printQueue(&intQueue)

	fmt.Println("\nSwitching to New Implementation:")
	newIntQueue := Queue{}
	newIntQueue.add(42)
	newIntQueue.add(99)
	intQueue.changeImpl(&newIntQueue)

	fmt.Println("Original Queue after switch:")
	printQueue(&intQueue)

	fmt.Println("\n***************STRING***************")
	// Strings
	stringQueue := Queue{}
	stringQueue.add("U")
	stringQueue.add("smell")
	stringQueue.add("bad")

	fmt.Println("Original Queue:")
	printQueue(&stringQueue)

	fmt.Println("\nSwitching to New Implementation:")
	newStringQueue := Queue{}
	newStringQueue.add("This")
	newStringQueue.add("is")
	newStringQueue.add("new")
	stringQueue.changeImpl(&newStringQueue)

	fmt.Println("Original Queue after switch:")
	printQueue(&stringQueue)
}

func printQueue(q *Queue) {
	fmt.Printf("Size: %d, IsEmpty: %t\n", q.size(), q.isEmpty())
	for !q.isEmpty() {
		fmt.Println(q.remove())
	}
	fmt.Println()
}
