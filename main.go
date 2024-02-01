package main

import (
	"fmt"
)

func main() {
    algoFileList := [3]string{"bubblesort", "quicksort", "stack"}
    checkFiles(algoFileList)
    
    fmt.Println("\nWelcome to Alta! It is a program to help you learn your Algorithms and Data Structures.") 
    fmt.Println("This program will create files in the DataAlgoPractice folder that you can edit and practice with.")

    quit := false
    for !quit {
        fmt.Println("\nEnter in the number to choose the option")
        fmt.Println("0.Close Program\n1. Reset File\n2. Show Answer\n3. Test File")
        
        userInput := getUserInput(2)

        if userInput == 0 {
            quit = true
        }
        if userInput == 1 {
            resetMenu(algoFileList)
        }
        if userInput == 2 {
            showAnswerMenu()
        }
    }
}

func resetMenu(algoFileList [3]string) {
    quit := false
    for !quit {
        fmt.Println("\nSelect which file you want to reset:")
        fmt.Println("0.Exit Menu\n1. Bubble Sort\n2. Quick Sort\n3. Stack")
        userInput := getUserInput(2)
        if (userInput == 0) {
            quit = true
        } else if (userInput == 1) {
            prepareFile("bubblesort")
        } else if (userInput == 2) {
            prepareFile("quicksort")
        } else if (userInput == 3) {
            prepareFile("stack")
        }
    }
}

func showAnswerMenu() {
    quit := false
    for !quit {
        fmt.Println("\nSelect which answer you want to see:")
        fmt.Println("0.Exit Menu\n1. Bubble Sort\n2. Quick Sort\n3. Stack")
        userInput := getUserInput(2)
        
        if (userInput == 0) {
            quit = true
        }else if (userInput == 1) {
            showFile("bubblesort")
        } else if (userInput == 2) {
            showFile("quicksort")
        }
    }
}

