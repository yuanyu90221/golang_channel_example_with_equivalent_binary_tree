package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t1 *tree.Tree, ch chan int) {
	// close channel before leave
	defer close(ch)
	// define inner function
	var walk func(t1 *tree.Tree)
	walk = func(t1 *tree.Tree) {
		if t1 == nil {
			return
		}
		// walk Left child
		walk(t1.Left)
		ch <- t1.Value
		// walk Right child
		walk(t1.Right)
	}
	walk(t1)
}
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 { // ch1 or ch2 close
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}
func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
