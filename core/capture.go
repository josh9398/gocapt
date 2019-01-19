package core

import (
    "fmt"
    "log"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
)

func Capture(opts Options) {
    inactive, err := pcap.NewInactiveHandle(*opts.InterfaceName)
    if err != nil {
        log.Fatal(err)
    }
    defer inactive.CleanUp()

    // Call various functions on inactive to set it up the way you'd like:
    // Timeout: Capture packets in time blocks
    // Promiscuous: Capture every packet in its entirety
    // SnapLength: Number of bytes to read per packet
    if err = inactive.SetTimeout(pcap.BlockForever); err != nil {
        log.Fatalln("unable to set timeout", err)
    } else if err = inactive.SetPromisc(*opts.Promiscuous); err != nil {
        log.Fatalln("could not set promisc mode:", err)
    } else if err = inactive.SetSnapLen(*opts.SnapLength); err != nil {
        log.Fatalln("could not set snap length", err)
    }

    // Finally, create the actual handle by calling Activate:
    handle, err := inactive.Activate() // after this, inactive is no longer valid
    if err != nil {
        log.Fatalln("PCAP Activate error:", err)
    }

    // Now use your handle as you see fit.
    defer handle.Close()

    var dec gopacket.Decoder
    var ok bool
    if dec, ok = gopacket.DecodersByLayerName[*opts.Decoder]; !ok {
        log.Fatalln("No decoder named", *opts.Decoder)
    }

    source := gopacket.NewPacketSource(handle, dec)
    source.NoCopy = true
    source.DecodeStreamsAsDatagrams = true

    // Begin capture
    for packet := range source.Packets() {
        fmt.Println(packet)
    }
}
