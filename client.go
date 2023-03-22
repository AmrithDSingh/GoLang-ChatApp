package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    fmt.Println("Connected to server. Type 'quit' to exit.")

    scanner := bufio.NewScanner(os.Stdin)
    responseScanner := bufio.NewScanner(conn)

    for {
        fmt.Print("Enter a question (or 'quit' to exit): ")
        scanner.Scan()
        question := scanner.Text()
        if question == "quit" {
            break
        }

        fmt.Fprintln(conn, question)

        fmt.Print("Enter the answer: ")
        scanner.Scan()
        answer := scanner.Text()

        fmt.Fprintln(conn, answer)

        // Read the server's response and display the questions and answers
        fmt.Println("Here are the current questions and answers:")
        for responseScanner.Scan() {
            line := responseScanner.Text()
            if line == "" {
                break
            }
            fmt.Println(line)
        }
    }
}

