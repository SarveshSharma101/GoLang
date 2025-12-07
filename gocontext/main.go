package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// ContextWithTimeOut()
	// ContextWithCancel()
	// ContextWithDeadline()
	ContextForKillSignal()
}
func ContextForKillSignal() {

	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go DoSomeWork(ctx)
	<-sigChan
	cancel()

}

func ContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go DoSomeWork(ctx)
	time.Sleep(3 * time.Second)
	cancel()

}

func ContextWithTimeOut() {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	DoSomeWork(ctx)
}

func ContextWithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(11*time.Second))
	defer cancel()
	DoSomeWork(ctx)
}

func DoSomeWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled, stopping the work and signing off")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("Doing work")
		}
	}
}
