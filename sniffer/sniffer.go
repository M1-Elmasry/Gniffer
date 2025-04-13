package sniffer

import (
	"gniffer/logger"
	"log"
	"slices"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func StartSniffing(iface string, watchPorts []int, logFile string) {

	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		process(packet, watchPorts)
	}
}

func process(packet gopacket.Packet, ports []int) {
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)

	if tcpLayer != nil && ipLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		ip, _ := ipLayer.(*layers.IPv4)

		if slices.Contains(ports, int(tcp.DstPort)) {
			logger.LogConnection(ip.SrcIP.String(), ip.DstIP.String(), int(tcp.DstPort), tcp.SYN, tcp.ACK)
		}
	}
}
