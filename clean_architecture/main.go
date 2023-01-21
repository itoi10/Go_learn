package main

import (
	"context"
	"fmt"
	"os"

	"example.com/drivers"
)

func main() {
	ctx := context.Background()
	userDriver, err := drivers.InitializeUserDriver(ctx)
	if err != nil {
		fmt.Printf("failed to create UserDriver: %v\n", err)
		os.Exit(2)
	}

	userDriver.ServeUsers(ctx, ":8080")
}
