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

	var wg sync.WaitGroup

	ch := make(chan interface{})
	repository := repository.New()
	repository.Restore()

	initialTaskLen := len(repository.GetTasks())
	initialUserLen := len(repository.GetUsers())

	wg.Add(2)
	go service.Add(ctx, &wg, ch, &repository)
	go logger.LogChanges(ctx, &repository, initialTaskLen, initialUserLen)

	go func() {
		defer wg.Done()
		i := len(repository.Tasks)
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

	wg.Wait()
}
