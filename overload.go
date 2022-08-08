package main

import (
	"fmt"
	"log"
	"strconv"
)

func overload(allServer []*Server, m *int, t *int) {
	// 調査したipの配列
	searchedIp := make([]string, 0)
	for i := len(allServer); i > 0; i-- {
		// 調査済みのipは除く
		if includeIp(searchedIp, allServer[i - 1].ip) {
			continue
		}
		// ipが何回出現するかカウント
		count := 0
		// 応答時間
		totalResponseTime := 0
		// m番目に出てくるipまで計算
		for j, _ := range allServer {
			if allServer[i - 1].ip == allServer[i - j - 1].ip {
				// pingが反応していない場合は別でipが出力されるので応答時間は表示しない
				if allServer[i - j - 1].responseTime == "-" {
					count = 0
					totalResponseTime = 0
					break
				}
				count++
				// 応答時間の計算
				time, err := strconv.Atoi(allServer[i - j - 1].responseTime)
				if err != nil {
					log.Fatal(err)
				}
				totalResponseTime += time
				if count == *m {
					// 平均応答時間の計算
					// HACK: 小数点は切り捨てられる
					averageResponseTime := totalResponseTime / *m
					// 平均応答時間が超えていれば過負荷状態
					if averageResponseTime > *t {
						fmt.Println("ip:", allServer[i - 1].ip , "は過負荷状態です。")
						// 過負荷状態の期間を出力
						overloadStartTime, err := strconv.Atoi(allServer[i - 1].responseTime)
						if err != nil {
							log.Fatal(err)
						}
						overloadEndTime, err := strconv.Atoi(allServer[i - j - 1].responseTime)
						if err != nil {
							log.Fatal(err)
						}
						overloadPeriod := overloadStartTime - overloadEndTime
						fmt.Println("過負荷状態の期間は", overloadPeriod, "msです。")
					}
					count = 0
					totalResponseTime = 0
					break
				}
			}
		}
		// 調査したipの置き場
		searchedIp = append(searchedIp, allServer[i - 1].ip)
	} 
	fmt.Println("------------------------")
}