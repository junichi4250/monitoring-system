package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func pingTimeout(breakServer []*Server, allServer []*ServerResponse, N *int) (responseAllServer []*ServerResponse) {
	data, _ :=  os.Open("access.log")
	defer data.Close()
	
	scanner := bufio.NewScanner(data)

	for scanner.Scan(){
		if strings.HasSuffix(scanner.Text(), "-") {
			breakServer = append(breakServer, &Server{
				strings.Split(scanner.Text(), ",")[0],
				strings.Split(scanner.Text(), ",")[1],
				strings.Split(scanner.Text(), ",")[2],
				0,
				0,
			})
		}

		// pingを送るサーバーが故障していたものかチェック
		index := include(breakServer, strings.Split(scanner.Text(), ",")[1])
		if index != -1 {
			if !strings.HasSuffix(scanner.Text(), "-") {
				// 故障が改善していれば故障期間を測定
				recordStartTime, err := strconv.Atoi(breakServer[index].recordTime)
				if err != nil {
					log.Fatal(err)
				}
				recordEndTime, err := strconv.Atoi(strings.Split(scanner.Text(), ",")[0])
				if err != nil {
					log.Fatal(err)
				}
				breakServer[index].breakTime = recordEndTime - recordStartTime
			} else {
				// 故障が改善していなければtimeout回数を増やす
				breakServer[index].timeoutCount++
				if breakServer[index].timeoutCount >= *N {
					// n回以上タイムアウトの時は故障
					fmt.Println("ipアドレス:", breakServer[index].ip, "は", *N, "回以上タイムアウトしました。")
				}
			}
		}

		allServer = append(allServer, &ServerResponse{
			strings.Split(scanner.Text(), ",")[1],
			strings.Split(scanner.Text(), ",")[2],
		})
	}

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

	return allServer
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