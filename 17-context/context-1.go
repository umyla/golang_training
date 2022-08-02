package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	PrintData(ctx, "hello from test func")
}
func PrintData(ctx context.Context, input string) {
	select {
	case <-ctx.Done():
		fmt.Println("timeout")
		fmt.Println(ctx.Err())
	case <-time.After(2 * time.Second):
		fmt.Println(input)
	}
}
