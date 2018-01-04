package main

import (
	"fmt"
	"github.com/majestrate/ircproxy/proxy"
	"net"
	"os"
	"time"
	"xd/lib/log"
	"xd/lib/network/i2p"
	"xd/lib/util"
)

func printUsage(cmd string) {
	fmt.Printf("usage: %s ircserverhere.i2p\n", cmd)
}

func main() {
	if len(os.Args) == 1 {
		printUsage(os.Args[0])
		return
	}
	ircAddr := "127.0.0.1:6667"

	samAddr := "127.0.0.1:7656"
	serverName := os.Args[1]

	var l net.Listener
	var err error

	l, err = net.Listen("tcp", ircAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	server := proxy.NewServer(serverName)

	session := i2p.NewSession(util.RandStr(7), samAddr, "", map[string]string{})
	for {
		log.Infof("connecting to i2p via %s", samAddr)
		err = session.Open()
		if err == nil {
			break
		} else {
			log.Errorf("failed to open connection to i2p: %s", err.Error())
			time.Sleep(time.Second)
		}
	}
	log.Info("Connected to i2p")
	server.Dial = session.Dial
	server.DCCListener = session
	log.Info("Running irc proxy")
	err = server.Serve(l)
	if err != nil {
		log.Errorf("failed to serve: %s", err.Error())
	}
}
