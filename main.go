package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"otus/internal/model/user"
	"otus/internal/repository"
	"otus/internal/service"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("Ending of app...")
		cancel()
	}()

	var pwg sync.WaitGroup // producer
	var cwg sync.WaitGroup //consumer

	done := make(chan struct{})
	ch := make(chan interface{})
	repository := repository.New()

	cwg.Add(1)
	go service.Add(ctx, &cwg, ch, done, &repository)

	pwg.Add(1)
	go func() {
		defer pwg.Done()

		select {
		case <-ctx.Done():
			fmt.Println("Goroutine 'Create' is done")
			return
		case <-time.After(3 * time.Second):
			obj, err := service.Create(1, "Open", "Todo", "Need to do", time.Now(), "Medium", 12)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}
			ch <- obj
		}
	}()

	pwg.Add(1)
	go func() {
		defer pwg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine 'Create' is done")
			return
		case <-time.After(5 * time.Second):
			obj, err := service.Create("Petr", user.Manager, 3)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}
			ch <- obj
		}
	}()

	pwg.Add(1)
	go func() {
		defer pwg.Done()

		select {
		case <-ctx.Done():
			fmt.Println("Goroutine 'Create' is done")
			return
		case <-time.After(7 * time.Second):
			obj, err := service.Create(2, "Closed", "Todo", "Need to do", time.Now(), "Low", 12)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}
			ch <- obj
		}
	}()

	go func() {
		pwg.Wait()
		done <- struct{}{}
	}()

	cwg.Wait()

	for _, value := range repository.Users {
		fmt.Println(value)
	}
	for _, value := range repository.Tasks {
		fmt.Println(value)
	}
}
