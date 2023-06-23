// main.go
package main

import "fmt"

func main() {
    message := hello()
    fmt.Println(message)
}

func hello() string {
    return "Hello, World!"
}