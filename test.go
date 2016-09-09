package main

import "sync"

var (
	once sync.Once
)

func main() {
	// data, _ := way.StructToJSON()
	// b := 1
	// a := string(b)
	// fmt.Printf("%s\n", data)
	// fmt.Printf("Hellworld %s", a)

	done := make(chan struct{})
	c := make(chan int)
	go func() {
		println("1111")
		defer close(done)
		for {
			x, ok := <-c
			if !ok {
				println("close")
				return
			}
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	println("xxxx")
	close(c)
	<-done
}
