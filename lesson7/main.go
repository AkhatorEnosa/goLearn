// Lesson 7: Key-Value Pairs (Maps)

package main

import "fmt"

func main() {
	counts := map[string]int{}
	var item string

	for {
		// ask for item
		fmt.Print("Enter item (or 'done' to finish): ")
		fmt.Scanln(&item)

		if item == "done" {
			for item, count := range counts {
				fmt.Printf("- %s: %d\n", item, count)
			}
			break
		} else {
			counts[item]++
		}
	}
}
