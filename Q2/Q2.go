//for mac users, you need to install mpg123 first via brew install mpg123
//for windows users, you need to install mpg123 first via https://www.mpg123.de/download.shtml

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	currentMusicPath string
	pauseChannel     = make(chan bool)
	stopChannel      = make(chan bool)
	quitChannel      = make(chan bool)
	playChannel      = make(chan string)
	mpg123Cmd        *exec.Cmd
	mpg123Stdin      io.Writer
)

var (
	pauseSignal    = "SIGSTOP"
	continueSignal = "SIGCONT"
	stopSignal     = "SIGINT"
)

func sendCommand(signalName string) {
	signalMap := map[string]os.Signal{
		pauseSignal:    syscall.SIGSTOP,
		continueSignal: syscall.SIGCONT,
		stopSignal:     syscall.SIGINT,
	}

	sig, ok := signalMap[signalName]
	if !ok {
		fmt.Printf("Unknown signal: %s\n", signalName)
		return
	}

	if mpg123Cmd != nil && mpg123Cmd.Process != nil {
		mpg123Cmd.Process.Signal(sig)
	}
}

func musicPlayer() {
	for {
		select {
		case musicPath := <-playChannel:
			if _, err := os.Stat(musicPath); os.IsNotExist(err) {
				fmt.Println("file does not exist.")
				continue
			}

			if currentMusicPath != "" {
				fmt.Printf("Stopping %s.\n", currentMusicPath)
				sendCommand(stopSignal)
				if mpg123Cmd != nil && mpg123Stdin != nil {
					if err := mpg123Cmd.Process.Kill(); err != nil {
						fmt.Println("Error stopping music:", err)
					}
					currentMusicPath = ""
				}
			}

			cmd := exec.Command("mpg123", musicPath)
			stdin, err := cmd.StdinPipe()
			if err != nil {
				fmt.Println("Error creating stdin pipe:", err)
				continue
			}
			cmd.Stdout = nil
			cmd.Stderr = nil

			err = cmd.Start()
			if err != nil {
				fmt.Println("Error playing music:", err)
				continue
			}
			mpg123Cmd = cmd
			mpg123Stdin = stdin
			currentMusicPath = musicPath

			// Wait for stop or quit signal
			for {
				select {
				case <-pauseChannel:
					fmt.Printf("Playing %s is paused.\n", currentMusicPath)
					sendCommand(pauseSignal)
				case <-stopChannel:
					fmt.Println("Stopping music...")
					sendCommand(stopSignal)
					return
				case <-quitChannel:
					return
				default:
					// Sleep for a short duration to avoid busy loop
					time.Sleep(100 * time.Millisecond)
				}
			}
		case <-quitChannel:
			fmt.Println("Quitting musicPlayer.")
			return
		}
	}
}

func controller() {
	fmt.Println("Welcome to Simple Music Player!")
	for {
		var command string
		fmt.Print("Enter ONLY the command: open, play, pause, continue, exit): ")
		fmt.Scanln(&command)
		command = strings.TrimSpace(command)
		switch command {
		case "open":
			var directory string
			fmt.Print("Enter directory path: ")
			fmt.Scanln(&directory)
			directory = strings.TrimSpace(directory)
			if _, err := os.Stat(directory); os.IsNotExist(err) {
				fmt.Println("Directory does not exist.")
			} else {
				os.Chdir(directory)
				fmt.Printf("The current working directory is changed to %s.\n", directory)
			}
		case "play":
			var path string
			fmt.Print("Enter music file path: ")
			fmt.Scanln(&path)
			path = strings.TrimSpace(path)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				fmt.Println("Music file does not exist.")
			} else {
				playChannel <- path
			}
		case "pause":
			pauseChannel <- true
		case "continue":
			pauseChannel <- false
		case "exit":
			fmt.Println("Thanks for listening to music with us, goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		controller()
	}()
	go func() {
		defer wg.Done()
		musicPlayer()
	}()
	wg.Wait()
}
