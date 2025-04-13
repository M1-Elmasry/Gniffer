package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logFile *os.File

func Init(filePath string) {
	var err error
	logFile, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
}

func LogConnection(srcIP string, dstIP string, dstPort int, syn bool, ack bool) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logLine := fmt.Sprintf("[%s] %s ➜ %s:%d | SYN=%t ACK=%t\n", timestamp, srcIP, dstIP, dstPort, syn, ack)

	fmt.Print(logLine)

	if logFile != nil {
		_, _ = logFile.WriteString(logLine)
	}
}
