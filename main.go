package main

import (
	"fmt"
	"otus/internal/model/user"
	"otus/internal/repository"
	"otus/internal/service"
	"sync"
	"time"
)

func main() {

	var pwg sync.WaitGroup // producer
	var cwg sync.WaitGroup //consumer
	ch := make(chan interface{})
	repository := repository.New()

	cwg.Add(1)
	go repository.Save(&cwg, ch)

	pwg.Add(1)
	go func() {
		defer pwg.Done()
		obj, err := service.Create(1, "Open", "Todo", "Need to do", time.Now(), "Medium", 12)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		ch <- obj

	}()

	pwg.Add(1)
	go func() {
		defer pwg.Done()
		obj, err := service.Create("Petr", user.Manager, 3)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		ch <- obj

	}()

	pwg.Add(1)
	go func() {
		defer pwg.Done()
		obj, err := service.Create(2, "Closed", "Todo", "Need to do", time.Now(), "Low", 12)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		ch <- obj

	}()

	go func() {
		pwg.Wait()
		close(ch)
	}()

	cwg.Wait()

	for _, value := range repository.Users {
		fmt.Println(value)
	}
	for _, value := range repository.Tasks {
		fmt.Println(value)
	}
}
