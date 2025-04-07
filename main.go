package main

import (
	"fmt"
	"otus/internal/model/user"
	"otus/internal/repository"
	"otus/internal/service"
	"time"
)

func main() {

	repository := repository.New()
	err := service.Create(&repository, 1, "Open", "Todo", "Need to do", time.Now(), "Medium", 12)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	time.Sleep(1 * time.Second)
	err = service.Create(&repository, "Petr", user.Manager, 3)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	time.Sleep(1 * time.Second)
	err = service.Create(&repository, 2, "Closed", "Todo", "Need to do", time.Now(), "Low", 12)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	for _, value := range repository.Users {
		fmt.Println(value)
	}
	for _, value := range repository.Tasks {
		fmt.Println(value)
	}
}
