package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Server struct {
	recordTime string
	ip string
	responseTime string
	breakTime int
	timeoutCount int
}

func main() {
	var N = flag.Int("timeout", 5, "timeout flag")
	flag.Parse()

	data, _ :=  os.Open("access.log")
	defer data.Close()
	scanner := bufio.NewScanner(data)

	server := []*Server{}

	for scanner.Scan(){
		if strings.HasSuffix(scanner.Text(), "-") {
			server = append(server, &Server{
				strings.Split(scanner.Text(), ",")[0],
				strings.Split(scanner.Text(), ",")[1],
				strings.Split(scanner.Text(), ",")[2],
				0,
				0,
			})
		}

		// pingを送るサーバーが故障していたものかチェック
		index := include(server, strings.Split(scanner.Text(), ",")[1])
		if index != -1 {
			if !strings.HasSuffix(scanner.Text(), "-") {
				// 故障が改善していれば故障期間を測定
				recordStartTime, err := strconv.Atoi(server[index].recordTime)
				if err != nil {
					log.Fatal(err)
				}
				recordEndTime, err := strconv.Atoi(strings.Split(scanner.Text(), ",")[0])
				if err != nil {
					log.Fatal(err)
				}
				server[index].breakTime = recordEndTime - recordStartTime
			} else {
				// 故障が改善していなければtimeout回数を増やす
				server[index].timeoutCount++
				if server[index].timeoutCount >= *N {
					// n回以上タイムアウトの時は故障
					fmt.Println(*N, "回以上タイムアウトしました。")
					fmt.Println("ipアドレス:", server[index].ip)
				}
			}
		}
	}

	for _, v := range server {
		// 故障期間が0になる場合は続けて2回以上timeoutした時なので表示からは除く
		if v.breakTime == 0 {
			continue
		}
		// 故障しているサーバーのipアドレス表示
		fmt.Println("故障サーバーip:" , v.ip)
		// 故障期間の出力
		fmt.Println("故障期間", v.breakTime)
	}
}

// 配列に指定した文字が何番目に含まれているか
func include(server []*Server, target string) int {
	for num, v := range server {
		// すでに設定されているものは除く
		if v.breakTime != 0 {
			continue
		}
		if v.ip == target {
			return num
		}
	}
	return -1
}