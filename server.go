package main

import (
    "bufio"
    "fmt"
    "net"
)

func servermain() {
    listener, err := net.Listen("tcp", ":8000")
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    fmt.Println("Server started. Listening for connections on port 8000...")

    qaMap := make(map[string]string)

    for {
        conn, err := listener.Accept()
        if err != nil {
            panic(err)
        }

        go handleConnection(conn, qaMap)
    }
}

func handleConnection(conn net.Conn, qaMap map[string]string) {
    defer conn.Close()

    scanner := bufio.NewScanner(conn)

    for {
        fmt.Fprint(conn, "Enter a question (or 'quit' to exit): ")
        scanner.Scan()
        question := scanner.Text()
        if question == "quit" {
            break
        }

        fmt.Fprint(conn, "Enter the answer: ")
        scanner.Scan()
        answer := scanner.Text()

        qaMap[question] = answer
    }

    fmt.Fprintln(conn, "Here are the current questions and answers:")
    for q, a := range qaMap {
        fmt.Fprintf(conn, "Q: %s\nA: %s\n\n", q, a)
    }
}
