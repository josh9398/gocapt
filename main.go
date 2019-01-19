package main

import (
    "log"
    "./core"
)

func main() {
    opts, err := core.GetOptions()
    if err != nil {
        log.Fatal(err)
    }

    // Create capture handler
    core.Capture(opts)
}
