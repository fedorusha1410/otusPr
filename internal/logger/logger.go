package logger

import (
	"context"
	"fmt"
	"log"
	"otus/internal/repository"
	"time"
)

func LogChanges(ctx context.Context, repo *repository.Repository) {
	var prevTaskLen, prevUserLen int
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine 'LogChanges' is done")
			return
		case <-ticker.C:
			tasks := repo.GetTasks()
			users := repo.GetUsers()

			if len(tasks) > prevTaskLen {
				lastTask := tasks[len(tasks)-1]
				log.Println("New tasks:")
				log.Printf("Task ID: %d\t, task name: %s ", lastTask.GetId(), lastTask.Title)
				prevTaskLen = len(tasks)
			}

			if len(users) > prevUserLen {
				lastUser := users[len(users)-1]
				log.Println("New users:")
				log.Printf("User ID: %d\t user Name: %s", lastUser.GetId(), lastUser.Name)
				prevUserLen = len(users)
			}
		}

	}
}
