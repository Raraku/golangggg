package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// values1 := make([]int, 0)
	// values2 := make([]int, 0)
	checker := make(map[int]int)
	ch1 := make(chan int)
	ch2 := make(chan int)
	quit := make(chan int, 2)
	queue := make([]int, 0, 2)
	//Walk 1

	go func() {
		Walk(t1, ch1)
		close(ch1)
		quit <- 1
	}()
	for i := range ch1 {
		_, ok := checker[i]
		if ok {
			checker[i] = checker[i] + 1
		} else {
			checker[i] = 1
		}
	}
	go func() {
		Walk(t2, ch2)
		close(ch2)
		quit <- 1
	}()
	for i := range ch2 {
		_, ok := checker[i]
		if ok {
			checker[i] = checker[i] + 1
		} else {
			checker[i] = 1
		}
	}
	for {
		select {
		case <-quit:
			if len(queue) == 0 {
				queue = append(queue, 1)
				break
			} else {
				for _, element := range checker {
					if element == 1 {
						return false
					}
				}
				return true

			}
		}
	}
}

func main() {
	// values := make([]int,0)
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
