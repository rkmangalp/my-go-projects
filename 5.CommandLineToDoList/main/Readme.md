# Command Line To-Do List

This is a simple command-line To-Do List application written in Go. The application allows users to add, remove, complete, and list tasks through a command-line interface.

## Features

- Add new tasks
- Remove tasks by number
- Mark tasks as complete
- List all tasks

## Requirements

- Go 1.16 or later

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/rkmangalp/my-go-projects.git
    cd CommandLineToDoList
    ```

2. **Initialize a new Go module:**

    ```sh
    go mod init https://github.com/rkmangalp/my-go-projects/CommandLineToDoList
    ```

3. **Download dependencies:**

    ```sh
    go mod tidy
    ```

## Usage

1. **Build the application:**

    ```sh
    go build
    ```

2. **Run the application:**

    ```sh
    ./CommandLineToDoList
    ```

3. **Commands:**

    - `add` : Add a new task.
    - `remove` : Remove a task by its number.
    - `complete` : Mark a task as complete by its number.
    - `list` : List all tasks.
    - `exit` : Exit the application.

### Example

```sh
Enter command (add, remove, complete, list, or exit):
add
Enter task to add:
Write a README file
Task added successfully!

Enter command (add, remove, complete, list, or exit):
list
Your tasks:
1. Write a README file

Enter command (add, remove, complete, list, or exit):
complete
Enter task number to mark complete:
1
Task 'Write a README file' marked complete.

Enter command (add, remove, complete, list, or exit):
list
Your tasks:
1. **COMPLETED** - Write a README file

Enter command (add, remove, complete, list, or exit):
exit
Exiting...
