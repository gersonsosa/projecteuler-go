package main

import (
    "fmt"
    "github.com/gersonsosa/projecteuler-go/datastructures"
)

func main() {
    maxHeap := datastructures.NewMaxHeap()
    maxHeap.Add(100)
    pollResult, _ := maxHeap.Poll()
    fmt.Printf("Pop result %v\n", pollResult)
    fmt.Println("Sucess!");
}
