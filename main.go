package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

var onlyOneSignalHandler = make(chan struct{})
var shutdownHandler chan os.Signal

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
// Only one of SetupSignalContext and SetupSignalHandler should be called, and only can
// be called once.
func SetupSignalHandler() <-chan struct{} {
	return SetupSignalContext().Done()
}

// SetupSignalContext is same as SetupSignalHandler, but a context.Context is returned.
// Only one of SetupSignalContext and SetupSignalHandler should be called, and only can
// be called once.
func SetupSignalContext() context.Context {
	close(onlyOneSignalHandler) // panics when called twice

	shutdownHandler = make(chan os.Signal, 2)

	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(shutdownHandler, shutdownSignals...)
	go func() {
		<-shutdownHandler
		cancel()
		<-shutdownHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	return ctx
}

// RequestShutdown emulates a received event that is considered as shutdown signal (SIGTERM/SIGINT)
// This returns whether a handler was notified
func RequestShutdown() bool {
	if shutdownHandler != nil {
		select {
		case shutdownHandler <- shutdownSignals[0]:
			return true
		default:
		}
	}

	return false
}

// initial implementation
func main_does_not_forward_SIGTERM_to_bash() {
	cmd := exec.Command("./script.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func main() {
	cmd := exec.Command("./script.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	shutdownHandler = make(chan os.Signal, 2)
	signal.Notify(shutdownHandler, shutdownSignals...)
	go func() {
		<-shutdownHandler
		fmt.Println("sending SIGTERM to", cmd.Process.Pid)
		cmd.Process.Signal(syscall.SIGTERM)
		<-shutdownHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
