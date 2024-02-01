package main

import (
    "fmt"
    "log"
    "os"
    "strings"
	"strconv"
)

func flushStdinBuf() {
    var buf [1]byte
    for {
        length, err := os.Stdin.Read(buf[:])
        if length > 0 && buf[0] == '\n' || err != nil {
            break
        }
    }
}

func sliceOffNull(byteBuf []byte) {
    for i:=0; i<len(byteBuf); i++ {
        if byteBuf[i] == 0 {
            byteBuf = byteBuf[:i]
            break
        }
    }
}

func validInputCheck(inputBytes []byte, inputInt int, validInputSize int) bool{
    const maxInput int = 3
    const minInput int = 0

    emptyInput := inputBytes[0] == '\n'
    inRange := inputInt >= minInput && inputInt <= maxInput

    if inRange && !emptyInput {
        return true
    } 

    inputTooLarge := inputBytes[validInputSize-1] != '\n'
    if inputTooLarge {
        flushStdinBuf()
    } 

    fmt.Println("Invalid input. Please try again:")
    return false
}

func getUserInput(inputSize int) int{
    var isValid bool = false
    
    inputBytes := make([]byte, inputSize)
    var userInput int

    for isValid == false {
        _, err := os.Stdin.Read(inputBytes[:]) 
        if err != nil {
            log.Fatalln("Failed: Reading User Input\n", err)
        }

        sliceOffNull(inputBytes)

        stringInput := strings.TrimSpace(string(inputBytes[:inputSize]))

        inputInt, inputErr := strconv.Atoi(stringInput)
        if inputErr != nil {
            fmt.Println("input failed to convert to string")
        }
            
        isValid = validInputCheck(inputBytes, inputInt, inputSize)

        if isValid {
            userInput = inputInt
        }
    }

    return userInput
}
