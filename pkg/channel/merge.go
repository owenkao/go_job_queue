package channel

import (
	"fmt"
	"sync"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go Counter(chan1)
	go Squarer(chan2, chan1)
	Printer(chan2)
}

// Counter is number input
func Counter(out chan<- int) {
	for i := 1; i < 10; i++ {
		out <- i
	}
	close(out)
}

// Squarer is number square
func Squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

// Printer is print
func Printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 为每一个输入channel cs 创建一个 goroutine output
	// output 将数据从 c 拷贝到 out，直到 c 关闭，然后 调用 wg.Done
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// 启动一个 goroutine，用于所有 output goroutine结束时，关闭 out
	// 该goroutine 必须在 wg.Add 之后启动
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
