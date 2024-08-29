package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(5)
    counter := 0

    for i := 0; i < 5; i++ {
        go func() {
            counter++
            fmt.Println(i)
            wg.Done()
        }()
    }

    wg.Wait()
}

func MaxInt(a, b int) int {
    if a >= b {
        return a
    }

    return b
}
