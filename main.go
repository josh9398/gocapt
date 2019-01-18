package main

import (
    "fmt"
    "log"
    "./core"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
)

func main() {
    fmt.Println("Hello World")
    opts, err := core.GetOptions()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(*opts.InterfaceName)
	if handle, err := pcap.OpenLive(*opts.InterfaceName, 1500, true, pcap.BlockForever); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("tcp and port 80"); err != nil {  // optional
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			fmt.Println(packet)  // Do something with a packet here.
		}
	}
}
