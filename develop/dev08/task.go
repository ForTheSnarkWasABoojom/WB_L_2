package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Program started! Type 'exit' to complete.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Exiting shell...")
			break
		}

		processCommand(input)
	}
}

func processCommand(input string) {
	args := strings.Fields(input)

	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "cd":
		if len(args) > 1 {
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Error when changing directory:", err)
			}
		} else {
			fmt.Println("An argument for cd must be specified.")
		}

	case "pwd":
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error when getting the current directory:", err)
		} else {
			fmt.Println(currentDir)
		}

	case "echo":
		if len(args) > 1 {
			fmt.Println(strings.Join(args[1:], " "))
		} else {
			fmt.Println()
		}

	case "kill":
		if len(args) > 1 {
			pid := args[1]
			err := exec.Command("kill", pid).Run()
			if err != nil {
				fmt.Println("Error when ending the process:", err)
			}
		} else {
			fmt.Println("You must specify the process PID for kill.")
		}

	case "ps":
		cmd := exec.Command("ps", "aux")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error when running ps command:", err)
		}

	default:
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Unexpected error:", err)
		}
	}
}
