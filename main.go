package main

import (
    "fmt"
    "log"
    "./core"
)

func main() {
    fmt.Println("Hello World")
    opts, err := core.GetOptions()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(*opts.InterfaceName)
}
