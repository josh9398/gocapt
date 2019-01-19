package core

import "flag"

type Options struct {
    InterfaceName *string
    Decoder *string
    Promiscuous *bool
    SnapLength *int
}

func GetOptions() (Options, error) {
    o := Options{
        InterfaceName: flag.String("iface",  "","Network interface to bind to, if empty the default interface will be auto selected."),
        Decoder: flag.String("decoder", "Ethernet", "Name of the decoder to use"),
        Promiscuous: flag.Bool("promisc", true, "Set promiscuous mode"),
        SnapLength: flag.Int("snaplen", 1600, "Snap length (max number of bytes to read per packet)"),
    }

    flag.Parse()

    return o, nil
}
