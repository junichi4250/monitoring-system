package main

import (
	"flag"
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

	allServer = pingTimeout(breakServer, allServer, N)

	overload(allServer, m, t)
}