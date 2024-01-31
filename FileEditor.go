package main

import (
    "fmt"
    "os"
    "log"
)

func checkFiles(algorithms [3]string) string {
    var algo string
    for i:=0; i < len(algorithms); i++ {
        algo = algorithms[i] 

        filePath := fmt.Sprintf("./DataAlgoPractice/%s.go", algo)
        _, err := os.Stat(filePath) 

        if err != nil {
            openFile(algo)
        }
    }
    return "pass"
}

func openFile(algo string) {
    fileName := fmt.Sprintf("DataAlgoPractice/%s.go", algo)

    createdFile, err := os.Create(fileName)

    if err != nil {
        log.Fatalln("Failed: Creating File")
    }

    resetFile(createdFile, algo)

    defer createdFile.Close()
}

func resetFile(practiceFile *os.File, algo string) {
    fileSource := fmt.Sprintf("./DataAlgoSource/%sIncomplete.txt", algo)

    byteAlgoCode, sourceErr := os.ReadFile(fileSource)
    if sourceErr != nil {
        log.Fatalln("Failed: Reading Source File")
    }

    algoCode := string(byteAlgoCode)

    _, err := practiceFile.WriteString(algoCode)
    if err != nil {
        log.Fatalln("Failed: Resetting File")
    }
}

//func showFile(practiceFile
