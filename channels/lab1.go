package channels

import "fmt"

func Run() {
	ch := make(chan int)

	go func() {
		ch <- 2
	}()
	x := <-ch
	fmt.Print(x)
}
