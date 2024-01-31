package main

import (
	"fmt"
)

func resetMenu() {
    fmt.Println("\nSelect which file you want to reset:")
    userInput := getUserInput(2)
    fmt.Println(userInput)

}

func main() {
    algoFileList := [3]string{"bubblesort", "quicksort", "stack"}
    checkFiles(algoFileList)
    
    fmt.Println("\nWelcome to Alta! It is a program to help you learn your Algorithms and Data Structures.") 
    fmt.Println("This program will create files in the DataAlgoPractice folder that you can edit and practice with.")
    fmt.Println("\nEnter in the number to choose the option")
    fmt.Println("1. Reset File\n2. Show Answer\n3. Test File")
    
    userInput := getUserInput(2)

    if userInput == 1 {
        resetMenu()
    }
}
