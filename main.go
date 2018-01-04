package main

import (
	"fmt"
	"os"
)

func printUsage(cmd string) {
	fmt.Printf("usage: %s ircserverhere.i2p\n", cmd)
}

func main() {
	if len(os.Args) == 1 {
		printUsage(os.Args[0])
		return
	}
	server := os.Args[1]
}
