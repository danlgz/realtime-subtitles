package main

import (
	"context"
	"fmt"
	"sync"
)

// App struct
type App struct {
	ctx context.Context
	recordSignal chan int // 1: play, 0: stop
	isRecording bool
	wg sync.WaitGroup
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		recordSignal: make(chan int),
	}
}

func (a *App) recordingController() {
	for {
		select {
		case signal := <-a.recordSignal:
			if signal == 1 {
				if a.isRecording { continue }
				a.isRecording = true
				go func() {
					a.wg.Add(1)
					defer a.wg.Done()
					record(&a.recordSignal)
				}()

				fmt.Println("Start recording")
			} else {
				if !a.isRecording { continue }
				a.isRecording = false
				close(a.recordSignal)

				a.wg.Wait()
				a.recordSignal = make(chan int)

				fmt.Println("Stop recording")
			}
		}
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// lister record signal
	go a.recordingController()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) StartRecord() {
	a.recordSignal <- 1
}

func (a *App) StopRecord() {
	a.recordSignal <- 0
}
