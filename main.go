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

	// for _, v := range server {
	// 	// 故障期間が0になる場合は続けて2回以上timeoutした時なので表示からは除く
	// 	if v.breakTime == 0 {
	// 		continue
	// 	}
	// 	// 故障しているサーバーのipアドレス表示
	// 	fmt.Println("故障サーバーip:" , v.ip)
	// 	// 故障期間の出力
	// 	fmt.Println("故障期間", v.breakTime)
	// }

	breakDown(allIp, server, breakServer, N)

}
