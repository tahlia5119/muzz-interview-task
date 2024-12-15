package main

import (
	"context"
	"log"
	"muzz/service"
)

// main function to run the service
func main() {
	ctx := context.Background()

	config, err := service.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	s, err := service.NewService(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	defer s.Close()

	s.Start()

	s.Wait()
}
