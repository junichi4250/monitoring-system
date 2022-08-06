package main

import (
	"bufio"
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
}

func main() {
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
			})
		}

		// pingを送るサーバーが故障していたものかチェック
		index := include(server, strings.Split(scanner.Text(), ",")[1])
		if index != -1 && server[index].breakTime == 0 {
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
			}
		}
	}

	for _, v := range server {
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