package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"otus/internal/logger"
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

	ch := make(chan interface{})
	repository := repository.New()

	cwg.Add(1)
	go service.Add(ctx, &cwg, ch, &repository)
	go logger.LogChanges(ctx, &repository)

	pwg.Add(1)

	go func() {
		defer pwg.Done()
		i := 0
		for {
			i++
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 'Create' is done")
				return
			case <-time.After(3 * time.Second):
				obj, err := service.Create(i, "Closed", fmt.Sprintf("Test №%d", i), "Need to do", time.Now(), "Low", i)
				if err != nil {
					fmt.Printf("error: %v", err)
					return
				}
				ch <- obj
				obj, err = service.Create(fmt.Sprintf("Test №%d", i), user.Manager, i)

				if err != nil {
					fmt.Printf("error: %v", err)
					return
				}
				ch <- obj

			}

		}
	}()

	pwg.Wait()
	cwg.Wait()

}
