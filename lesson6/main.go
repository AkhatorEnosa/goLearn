// Lists of Data (Slices)

package main

import "fmt"

func main() {
	shoppingList := []string{"Milk", "Bread", "Eggs"}
	var newItem string
	
	for {
		for i, item := range shoppingList {
			fmt.Printf("%d. %s.\n", i+1, item)
		}

		fmt.Print("Add a new item: ")
		fmt.Scanln(&newItem)

		if newItem == "0" {
			fmt.Println("Exited the program")
			break
		}
		shoppingList = append(shoppingList, newItem)
		fmt.Printf("Added: %s\n", newItem)

		fmt.Printf("Your list contains %d items.\n", len(shoppingList))
	}
}


// idx := 1

// tasks = append(tasks[:idx], tasks[idx+1:]...)
// fmt.Println(tasks)