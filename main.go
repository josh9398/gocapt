package main

import (
    "fmt"
    "github.com:josh9398/gocapt/core"
)

func main() {
    fmt.Println("Hello World")
    opts, err := core.ParseOptions()
    if err != nil {
        return nil, err
    }
    if opts.InterfaceName == "" {
        fmt.Println("No arg")
    } else {
        fmt.Println(opts.InterfaceName)
    }
}
