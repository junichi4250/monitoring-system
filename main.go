package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		N = flag.Int("timeout", 5, "timeout flag") // timeout回数
		m = flag.Int("responseTimeCount", 2, "responseTimeCount flag") // 平均応答時間を計測する回数
		t = flag.Int("averageResponseTime", 10, "averageResponseTime flag") // 平均応答時間
	)
	flag.Parse()

	breakServer := []*Server{}
	allServer := []*ServerResponse{}

	breakServer, allServer = pingTimeout(breakServer, allServer, N)

	overload(allServer, m, t)

	for _, v := range breakServer {
		// 故障期間が0になる場合は続けて2回以上timeoutした時なので表示からは除く
		if v.breakTime == 0 {
			continue
		}
		// 故障しているサーバーのipアドレス表示
		fmt.Println("故障サーバーip:" , v.ip)
		// 故障期間の出力
		fmt.Println("故障期間", v.breakTime , "ms")
	}
}