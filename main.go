package main

import (
	"flag"
	"fmt"
	"gniffer/sniffer"
	"log"
	"strconv"
	"strings"
  "gniffer/logger"
)

func main() {
	portList := flag.String("ports", "", "Ports to sniff. Comma separated (ex. 80,443)")
	iface := flag.String("iface", "", "Interface to sniff (ex. eth0)")
	logFile := flag.String("log", "", "Output Log file path (ex. /var/log/gniffer.log)")

	flag.Parse()

	if *iface == "" || *portList == "" || *logFile == "" {
		log.Fatal("Missing arguments. you must to specify --iface, --ports and --log flags")
	}

	watchPorts := strings.Split(*portList, ",")
	intWatchPorts := make([]int, len(watchPorts))
	for i, port := range watchPorts {
		intPort, err := strconv.Atoi(port)
		if err != nil {
			log.Fatal(err)
		}
		if intPort < 1 || intPort > 65535 {
			log.Fatal("Invalid port number")
		}
		intWatchPorts[i] = intPort
	}

	fmt.Printf("sniffing on %s ports %s\n", *iface, strings.Join(watchPorts, ", "))


	logger.Init(*logFile)
	sniffer.StartSniffing(*iface, intWatchPorts, *logFile)

}
