package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("Example of context with value")

	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, "key1", "value1")
	ctxWithValue2 := context.WithValue(ctxWithValue, "key2", "value2")

	// for this context key2 will not be present
	printContext(ctxWithValue, "first context")
	// for this both key values will be present
	printContext(ctxWithValue2, "second context")

}

func printContext(ctx context.Context, s string) {

	fmt.Println("Context data for ", s)
	fmt.Println("key1", ctx.Value("key1"))
	fmt.Println("key2", ctx.Value("key2"))
}
