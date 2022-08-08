package main

import (
	"flag"
)

func main() {
	N := flag.Int("timeout", 5, "timeout flag") // timeout回数
	m := flag.Int("responseTimeCount", 2, "responseTimeCount flag") // 平均応答時間を計測する回数
	t := flag.Int("averageResponseTime", 10, "averageResponseTime flag") // 平均応答時間
	flag.Parse()

	allIp := []string{}
	server := []*Server{}
	breakServer := []*BreakServer{}
	timeoutServer := []*TimeoutServer{}
	allServer := []*Server{}
	subnetServer := []*SubnetServer{}
	breakSubnetServer := []*BreakSubnetServer{}
	
	allServer = breakDown(allIp, server, breakServer, timeoutServer, allServer, N)
	
	overload(allServer, m, t)

	subnet(allServer, subnetServer, breakSubnetServer, N)
}