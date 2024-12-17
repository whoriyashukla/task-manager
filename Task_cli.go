package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	id          int
	title       string
	description string
	status      TaskStatus
	createdAt   time.Time
	updatedAt   time.Time
}

type TaskStatus string

const (
	StatusTodo       TaskStatus = "TODO"
	StatusInProgress TaskStatus = "IN_PROGRESS"
	StatusCompleted  TaskStatus = "COMPLETED"
)

type TaskManager struct {
	tasks  []Task
	nextId int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  []Task{},
		nextId: 1,
	}
}

func (tm *TaskManager) AddTask(title, description string) Task {
	task := Task{
		id:          tm.nextId,
		title:       title,
		description: description,
		status:      StatusTodo,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}
	tm.tasks = append(tm.tasks, task)
	tm.nextId++
	return task
}
func (tm *TaskManager) ListTasks() []Task {
	return tm.tasks
}
func (tm *TaskManager) UpdateTaskStatus(id int, newStatus TaskStatus) error {
	for i, tasks := range tm.tasks {
		if tasks.id == id {
			tm.tasks[i].status = newStatus
			tm.tasks[i].updatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
func (tm *TaskManager) DeleteTask(id int) error {
	for i, tasks := range tm.tasks {
		if tasks.id == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

type CLI struct {
	taskManager *TaskManager
	reader      *bufio.Reader
}

func NewCLI() *CLI {
	return &CLI{
		taskManager: NewTaskManager(),
		reader:      bufio.NewReader(os.Stdin),
	}
}

func (cli *CLI) Run() {
	for {
		cli.displayMenu()
		choice := cli.readInput("Enter your choice: ")

		switch choice {
		case "1":
			cli.addTaskPrompt()
		case "2":
			cli.listTasksPrompt()
		case "3":
			cli.updateTaskStatusPrompt()
		case "4":
			cli.deleteTaskPrompt()
		case "5":
			fmt.Println("Exiting Task Tracker. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (cli *CLI) displayMenu() {
	fmt.Println("\n--- Task Tracker CLI ---")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Update Task Status")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
}
func (cli *CLI) readInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := cli.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (cli *CLI) addTaskPrompt() {
	title := cli.readInput("Enter task title: ")
	description := cli.readInput("Enter task description: ")
	task := cli.taskManager.AddTask(title, description)
	fmt.Printf("Task added successfully. id: %d\n", task.id)
}

func (cli *CLI) listTasksPrompt() {
	tasks := cli.taskManager.ListTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return

	}
	fmt.Println("--- Current Tasks ---")
	for _, task := range tasks {
		fmt.Printf("ID: %d | Title: %s | Status: %s | Created: %s\n",
			task.id, task.title, task.status, task.createdAt.Format("2006-01-02 15:04"))
	}
}

func (cli *CLI) updateTaskStatusPrompt() {
	idStr := cli.readInput("Enter task ID to update: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	fmt.Println("Select new status:")
	fmt.Println("1. TODO")
	fmt.Println("2. IN PROGRESS")
	fmt.Println("3. COMPLETED")
	statusChoice := cli.readInput("Enter status number: ")

	var newStatus TaskStatus
	switch statusChoice {
	case "1":
		newStatus = StatusTodo
	case "2":
		newStatus = StatusInProgress
	case "3":
		newStatus = StatusCompleted
	default:
		fmt.Println("Invalid status.")
		return
	}

	err = cli.taskManager.UpdateTaskStatus(id, newStatus)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Task status updated successfully.")
	}
}

func (cli *CLI) deleteTaskPrompt() {
	idStr := cli.readInput("Enter task ID to delete: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}

	err = cli.taskManager.DeleteTask(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Task deleted successfully.")
	}
}
func main() {
	cli := NewCLI()
	cli.Run()
}
