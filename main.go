package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	TaskName  string
	completed bool
}

var tasks []Task

func addTask(task string) {
	newTask := Task{TaskName: task, completed: false}
	tasks = append(tasks, newTask)
	fmt.Println("Task Added")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return
	}

	for i, task := range tasks {
		status := "n"
		if task.completed {
			status = "d"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.TaskName, status)
	}
}

func addComment(i int) {
	if i >= 1 && i <= len(tasks) {
		tasks[i-1].completed = true
		fmt.Println("Task marked as completed")
	} else {
		fmt.Println("invalid index")
	}
}

func editTask(i int, newString string) {
	if i >= 1 && i <= len(tasks) {
		tasks[i-1].TaskName = newString
		fmt.Println("Task edited")
	} else {
		fmt.Println("invalid index")
	}
}

func deleteTask(i int) {
	if i >= 1 && i <= len(tasks) {
		tasks = append(tasks[:i-1], tasks[i:]...)
		fmt.Println("Task deleted")
	} else {
		fmt.Println("invalid index")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nOptions")
		fmt.Println("1 Add Task")
		fmt.Println("2 List Tasks")
		fmt.Println("3 Mark as completed")
		fmt.Println("4 Edit Task")
		fmt.Println("5 Delete Task")
		fmt.Println("6 Exit")

		fmt.Print("Enter choice (1,2,3,4,5,6): ")

		if !scanner.Scan() {
			fmt.Println("\nInput closed")
			return
		}

		input := scanner.Text()
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter Task: ")
			if !scanner.Scan() {
				fmt.Println("\nInput closed")
				return
			}
			taskInput := scanner.Text()
			addTask(taskInput)

		case 2:
			listTasks()

		case 3:
			fmt.Print("Enter task index to mark completed: ")
			if !scanner.Scan() {
				fmt.Println("\nInput closed")
				return
			}
			indexInput, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid index")
				continue
			}
			addComment(indexInput)

		case 4:
			fmt.Print("Enter task index to edit: ")
			if !scanner.Scan() {
				fmt.Println("\nInput closed")
				return
			}
			indexInput, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid index")
				continue
			}

			fmt.Print("Enter new task name: ")
			if !scanner.Scan() {
				fmt.Println("\nInput closed")
				return
			}
			newTaskInput := scanner.Text()
			editTask(indexInput, newTaskInput)

		case 5:
			fmt.Print("Enter task index to delete: ")
			if !scanner.Scan() {
				fmt.Println("\nInput closed")
				return
			}
			indexInput, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid index")
				continue
			}
			deleteTask(indexInput)

		case 6:
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
