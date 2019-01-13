package core

import "fmt"

type Options struct {
    InterfaceName *string
}

func GetOptions() (Options, error) {
    o := Options{
        InterfaceName: flag.String("iface", "Network interface to bind to, if empty the default interface will be auto selected.")
    }

    flag.Parse

    return o, nil
}
