package main

import (
	"fmt"
	"strings"
)

func subnet(allServer []*Server, subnetServer []*SubnetServer, breakSubnetServer []*BreakSubnetServer, N *int) {
	// 全サブネットの取得
	for _, v := range allServer {
		// HACK: 今のままだと8の倍数のサブネットマスクしか対応できない
		subnet := v.ip[:strings.Index(v.ip, "/")]
		switch v.ip[strings.Index(v.ip, "/") + 1:] {
		case "8":
			subnet = v.ip[:strings.Index(v.ip, ".")]
		case "16":
			tmpSubnet := subnet[:strings.LastIndex(v.ip, ".")]
			subnet = tmpSubnet[:strings.LastIndex(v.ip, ".")]
		case "24":
			subnet = v.ip[:strings.LastIndex(v.ip, ".")]
		}

		if !includeSubnet(subnetServer, subnet) {
			subnetServer = append(subnetServer, &SubnetServer{
				subnet,
			})
		}
	}
	// サブネット内のサーバが全て故障（ping応答がすべてN回以上連続でタイムアウト）している場合はスイッチの故障とみなす
	// 開始から一致している文字列が連続でタイムアウトしているかを調べる
	timeoutCount := 0
	breakStartTime := "0"
	for _, v := range allServer {
		match, subnet := matchSubnet(subnetServer, v.ip)
		if !match {
			continue
		}
		// 故障している
		if v.responseTime == "-" {
			timeoutCount++
			// 故障開始期間
			if timeoutCount == 1 {
				breakStartTime = v.recordTime
			}
			// N回連続で故障していれば、スイッチが故障している
			if timeoutCount >= *N {
				// 故障サブネットを登録
				breakSubnetServer = append(breakSubnetServer, &BreakSubnetServer{
					subnet,
					breakStartTime,
					v.recordTime,
				})
				// カウントリセット
				timeoutCount = 0
			}
			//故障していない
		} else {
			timeoutCount = 0
		}
	}
	// 故障サブネットを出力
	for _, v := range subnetServer {
		printBreakSubnet(breakSubnetServer, v.subnet)
	}
}

func includeSubnet(subnetServer []*SubnetServer, subnet string) bool {
	for _, v := range subnetServer {
		if v.subnet == subnet {
			return true
		}
	}
	return false
}

func matchSubnet(subnetServer []*SubnetServer, ip string) (bool, string) {
	for _, v := range subnetServer {
		if strings.Index(ip, v.subnet) == 0 {
			// matchしている
			return true, v.subnet
		}
	}
	return false, "0"
}

func printBreakSubnet(breakSubnetServer []*BreakSubnetServer, ip string) {
	// 出力フォーマットを整えるためのカウント
	count := 0
	for _, v := range breakSubnetServer {
		if ip == v.subnet && count == 0 {
			fmt.Println("故障サブネット:")
			displaySubnet := v.subnet
			switch strings.Count(v.subnet, ".") {
			case 1:
				displaySubnet = v.subnet + ".0.0"
			case 2:
				displaySubnet = v.subnet + ".0"
			}
			fmt.Println(displaySubnet)
			count++
		}
		if ip == v.subnet && count > 0 {
			fmt.Println(v.breakStartTime, "~", v.breakEndTime)
		}
	}
}