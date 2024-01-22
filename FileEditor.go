package main

import (
    "fmt"
    "os"
    "log"
)

func createFiles(fileName string) {
    createdFile, err := os.Create(fileName)

    if err != nil {
        log.Fatalln("Failed: Creating File")
    }

    defer createdFile.Close()
}

func checkFiles(algorithms [2]string) string {
    var algo string
    for i:=0; i < len(algorithms); i++ {
        algo = algorithms[i] 
        filePath := fmt.Sprintf("./DataAlgoPractice/%s.go", algo)
        _, err := os.Stat(filePath) 

        if err != nil {
            fileName := filePath[2:]
            createFiles(fileName)
        }
    }

    return "pass"
}
