// Lesson 8 — Functions

package main

import "fmt"

func formatTask(index int, title string, completed bool) string {
	var formatted string
	if !completed {
		formatted = fmt.Sprintf("%d. %s []", index+1, title)
	} else {
		formatted = fmt.Sprintf("%d. %s [x]", index+1, title)
	}

	return formatted
}

func canCompleteTask(taskNum, totalTasks int) bool {
	return 1 <= taskNum && taskNum <= totalTasks
}

func main() {
	result1 := formatTask(0, "Learn Python", true)
	result2 := canCompleteTask(1, 5)

	fmt.Println(result1)
	fmt.Println(result2)
}