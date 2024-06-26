package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Create a new scanner that reads from standard input (keyboard)
	scanner := bufio.NewScanner(os.Stdin)

	// Initialize an empty slice to store tasks
	tasks := []string{}

	// Infinite loop to keep the program running until user chooses to exit
	for {
		// Prompt the user to enter a command
		fmt.Println(" Enter a command (add, remove, complete, list or exit):")

		// Read the next line of input from the user
		userInput := scanner.Scan()
		if !userInput {
			break // Exit the loop if scanning fails (Ctrl+D for EOF)
		}

		// Retrieve the command entered by the user
		command := scanner.Text()

		// Process the command based on user input
		switch command {
		case "add":
			// Prompt user to enter a task to add
			fmt.Println("Enter a task to add:")

			// Read the next line of input (task description)
			scanner.Scan()
			task := scanner.Text()

			// Add the task to the tasks slice
			tasks = append(tasks, task)

			// Confirm to the user that the task was added successfully
			fmt.Println("Task was added successfully")

		case "remove":
			// Prompt user to enter a task to add
			fmt.Println("Enter task number to be remove:")

			// Read the next line of input (task number as string)
			scanner.Scan()
			taskNumberStr := scanner.Text()

			// Convert the task number from string to integer
			taskNumberInt, err := strconv.Atoi(taskNumberStr)

			//check for error
			if err != nil {
				fmt.Println("Invalid Inpur. Please enter a number.")
				continue
			}
			// Check if the task number is within valid range
			if taskNumberInt < 1 || taskNumberInt > len(tasks) {
				fmt.Println("Task Number is out of Range.")
				continue
			}

			tasks = append(tasks[:taskNumberInt-1], tasks[taskNumberInt:]...)

			// Confirm to the user that the task was removed successfully
			fmt.Println("Task removed successfully!")

		case "complete":
			// Prompt user to enter a task to add
			fmt.Println("Enter task number to Mark complete")

			// Read the next line of input (task number as string)
			scanner.Scan()
			taskNumberStr := scanner.Text()

			// Convert the task number from string to integer
			taskNumberInt, err := strconv.Atoi(taskNumberStr)
			//check for error
			if err != nil {
				fmt.Println("Invalid Input. Please enter a number")
			}
			// Check if the task number is within valid range
			if taskNumberInt < 1 || taskNumberInt > len(tasks) {
				fmt.Println("Task number is out of range")
			}
			// Mark the task as completed by updating its text
			fmt.Printf("Task %s is maked as Complete", tasks[taskNumberInt-1])
			tasks[taskNumberInt-1] = "**{COMPLETED} - " + tasks[taskNumberInt-1]

		case "list":
			// Check if there are any tasks in the list
			if len(tasks) == 0 {
				fmt.Println("You dont have any tasks yet.")
			} else {
				// Display all tasks in numbered format
				fmt.Println("Your tasks: ")
				for i, task := range tasks {
					fmt.Printf("%d. %s", i+1, task)
					fmt.Println()
				}
			}

		case "exit":
			//exit program gracefully
			fmt.Println("Exiting...")
			return

		default:
			//handle invalid command
			fmt.Println("Invalid command.")
		}
	}
}
