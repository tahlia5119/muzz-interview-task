package main

import (
	"context"
	"muzz/service"
)

// main function to run the service
func main() {
	ctx := context.Background()

	s := service.NewService(ctx)

	defer s.Close()

	s.Start()

	s.Wait()
}
