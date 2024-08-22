package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	n := 1
	go func() {
		for {
			time.Sleep(time.Second * 2)
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				time.Sleep(time.Second * 1)
				n++
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
