package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func breakDown(allIp []string, server []*Server, breakServer []*BreakServer, timeoutServer []*TimeoutServer, allServer []*Server, N *int) []*Server {
	data, _ :=  os.Open("access.log")
	defer data.Close()
	scanner := bufio.NewScanner(data)
	
	for scanner.Scan(){
		allServer = append(allServer, &Server{
			strings.Split(scanner.Text(), ",")[0],
			strings.Split(scanner.Text(), ",")[1],
			strings.Split(scanner.Text(), ",")[2],
			true,
			0,
		})

		// 含まれていなければipリストに追加
		if !includeIp(allIp, strings.Split(scanner.Text(), ",")[1]) {
			allIp = append(allIp, strings.Split(scanner.Text(), ",")[1])
		}

		// 失敗したサーバー
		if strings.HasSuffix(scanner.Text(), "-") {
			server = append(server, &Server{
				strings.Split(scanner.Text(), ",")[0],
				strings.Split(scanner.Text(), ",")[1],
				strings.Split(scanner.Text(), ",")[2],
				true,
				0,
			})
		}

		// pingを送るサーバーが故障していたものかチェック
		index := include(server, strings.Split(scanner.Text(), ",")[1])
		if index != -1 {
			if !strings.HasSuffix(scanner.Text(), "-") {
				// 故障から改善
				server[index].failuare = false
				// 故障が改善していれば故障期間を測定
				breakServer = append(breakServer, &BreakServer{
					strings.Split(scanner.Text(), ",")[1],
					server[index].recordTime,
					strings.Split(scanner.Text(), ",")[0],
				})
			} else {
				// 故障が改善していなければtimeout回数を増やす
				server[index].timeoutCount++
				if server[index].timeoutCount >= *N {
					// n回以上タイムアウトの時は故障
					if !includeIpTimeoutServer(timeoutServer, server[index].ip) {
						timeoutServer = append(timeoutServer, &TimeoutServer{
							server[index].ip,
						})
					}
				}
			}
		}
	}

	for _, ip := range allIp {
		// 故障しているipか判定
		if !includeIpBreakServer(breakServer, ip) {
			continue
		}
		// 故障しているサーバーのipアドレス表示
		fmt.Println("故障サーバーip:" , ip)
		// 故障期間の出力
		fmt.Println("故障期間:")
		// 故障サーバーのipに紐づく故障期間を出す
		for _, v := range breakServer {
			if v.ip == ip {
				fmt.Println(v.breakStartTime, "~", v.breakEndTime)
			}
		}
	}
	fmt.Println("------------------------")

	// timeout
	fmt.Println(*N, "回以上連続タイムアウト")
	fmt.Println("ipアドレス:")
	for _, v  := range timeoutServer {
		fmt.Println(v.ip)
	}
	fmt.Println("------------------------")
	return allServer
}

// 配列に指定した文字が含まれているか
func includeIp(ip []string, newIp string) bool {
	for _, v := range ip {
		if v == newIp {
			return true
		}
	}
	return false
}

// 配列に指定した文字が含まれているか
func includeIpBreakServer(breakServer []*BreakServer, ip string) bool {
	for _, v := range breakServer {
		if v.ip == ip {
			return true
		}
	}
	return false
}

// 配列に指定した文字が何番目に含まれているか
func include(server []*Server, target string) int {
	for num, v := range server {
		// すでに改善されているものは除く
		if !v.failuare {
			continue
		}
		if v.ip == target {
			return num
		}
	}
	return -1
}

// 配列に指定した文字が含まれているか
func includeIpTimeoutServer(timeoutServer []*TimeoutServer, ip string) bool {
	for _, v := range timeoutServer {
		if v.ip == ip {
			return true
		}
	}
	return false
}