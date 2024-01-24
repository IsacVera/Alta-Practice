package main

import (
    "fmt"
    "os"
    "log"
)

func resetFile(practiceFile *os.File, algo string, state string) {
    var fileSource string
    if state == "complete" {
        fileSource = fmt.Sprintf("./DataAlgoSource/%sComplete.txt", algo)
    } else {
        fileSource = fmt.Sprintf("./DataAlgoSource/%sIncomplete.txt", algo)
    }

    byteAlgoCode, sourceErr := os.ReadFile(fileSource)
    if sourceErr != nil {
        log.Fatalln("Failed: Reading Source File")
    }

    algoCode := string(byteAlgoCode)

    _, err := practiceFile.WriteString(algoCode)
    if err != nil {
        log.Fatalln("Failed: Reseting File")
    }
}

func createFile(fileName string, algo string) {
    createdFile, err := os.Create(fileName)

    if err != nil {
        log.Fatalln("Failed: Creating File")
    }

    resetFile(createdFile, algo, "reset")
    

    defer createdFile.Close()
}

func checkFiles(algorithms [3]string) string {
    var algo string
    for i:=0; i < len(algorithms); i++ {
        algo = algorithms[i] 

        filePath := fmt.Sprintf("./DataAlgoPractice/%s.go", algo)
        _, err := os.Stat(filePath) 

        if err != nil {
            fileName := filePath[2:]
            createFile(fileName, algo)
        }
    }

    return "pass"
}
