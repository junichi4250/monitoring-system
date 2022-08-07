package main

import (
	"flag"
)

func main() {
	N := flag.Int("timeout", 5, "timeout flag")
	flag.Parse()

	allIp := []string{}
	server := []*Server{}
	breakServer := []*BreakServer{}
	timeoutServer := []*TimeoutServer{}

	breakDown(allIp, server, breakServer, timeoutServer, N)

}
