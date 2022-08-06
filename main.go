package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ :=  os.Open("access.log")
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan(){
		if strings.HasSuffix(scanner.Text(), "-") {
			fmt.Println(scanner.Text())
		}
	}
}