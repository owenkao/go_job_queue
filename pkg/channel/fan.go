package channel

import "fmt"

// Start test
func Start() {
	in := producer(1, 2, 3, 4)
	ch := consumer(in)

	// consumer
	for ret := range ch {
		fmt.Printf("%3d", ret)
	}

	fmt.Println()
}

func producer(nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for n := range nums {
			fmt.Printf("%d\n", n)
			ch <- n
		}

	}()

	return ch
}

func consumer(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			fmt.Printf("!!! %d\n", n*n)
			out <- n * n
		}
	}()

	return out
}
