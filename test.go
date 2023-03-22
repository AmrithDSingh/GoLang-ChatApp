package main

import (
    "bufio"
    "fmt"
    "os"
)

func mai() {
    scanner := bufio.NewScanner(os.Stdin)
    qaMap := make(map[string]string)

    for {
        fmt.Print("Enter a question (or 'quit' to exit): ")
        scanner.Scan()
        question := scanner.Text()
        if question == "quit" {
            break
        }

        fmt.Print("Enter the answer: ")
        scanner.Scan()
        answer := scanner.Text()

        qaMap[question] = answer
    }

    fmt.Println("Here are the questions and answers:")
    for q, a := range qaMap {
        fmt.Printf("Q: %s\nA: %s\n\n", q, a)
    }
}
