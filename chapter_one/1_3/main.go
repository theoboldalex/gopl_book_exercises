package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    // print each cmdline arg using a loop
    startLoop := time.Now()

    for _, arg := range os.Args[1:] {
        fmt.Println(arg)
    }

    diffLoop := time.Since(startLoop).Seconds()

    fmt.Printf("Loop took %f seconds to execute.\n", diffLoop)


    // print each cmdline arg using strings.Join
    startJoin := time.Now()

    fmt.Println(strings.Join(os.Args[1:], " "))

    diffJoin := time.Since(startJoin).Seconds()

    fmt.Printf("Join took %f seconds to execute.\n", diffJoin)
}
