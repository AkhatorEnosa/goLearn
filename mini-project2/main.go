// Mini-Project 2: Interactive CLI To-Do List Manager

package main

import "fmt"

func main() {
	tasks := []string{}
	done := []bool{}
	var choice int

	for {
		fmt.Println("\n=== To-Do Manager ===")
		fmt.Println("1. Add task")
		fmt.Println("2. View tasks")
		fmt.Println("3. Complete task")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		var task string
		var taskNum int

		switch choice {
		case 1:
			// Add task
			fmt.Print("Add a task: ")
			fmt.Scanln(&task)
			tasks = append(tasks, task)
			done = append(done, false)
		case 2:
			// View tasks
			if len(tasks) > 0 {
				for i, task := range tasks {
					if done[i] {
						fmt.Printf("%d. %s [x]\n", i+1, task)
					} else {
						fmt.Printf("%d. %s [ ]\n", i+1, task)
					}
				}
			} else {
				fmt.Println("No tasks yet.\n")
			}
		case 3:
			// Complete task
			fmt.Println("Complete a task(Input the task number): ")
			fmt.Scanln(&taskNum)

			if 1 <= taskNum && taskNum <= len(tasks) {
				done[taskNum-1] = true
			} else {
				fmt.Println("Invalid task number.\n")
			}
		case 4:
			// Delete task
			fmt.Println("Delete a task(Input the task number): ")
			fmt.Scanln(&taskNum)
			if 1 <= taskNum && taskNum <= len(tasks) {
				tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)
				done = append(done[:taskNum-1], done[taskNum:]...)
			} else {
				fmt.Println("Invalid task number.\n")
			}
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Choose 1-5.")
		}
	}
}