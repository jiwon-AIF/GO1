package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


func main() {

	fmt.Println("숫자를 입력하세요")

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1, _ := strconv.ParseFloat(line, 64)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.ParseFloat(line, 64)

	fmt.Printf("입력하신 숫자는 %f, %f 입니다.", n1, n2)

	fmt.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+" {
		fmt.Printf("%.2f + %.2f = %.2f", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%.2f - %.2f = %.2f", n1, n2, n1-n2)
	} else if line == "*" {
		fmt.Printf("%.2f * %.2f = %.2f", n1, n2, n1*n2)
	} else if line == "/" {
		fmt.Printf("%.2f / %.2f = %.2f", n1, n2, n1/n2)
	} else if line == "q" {
		os.Exit(1)
	} else {
		fmt.Println("잘못 입력하셨습니다.")
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
time.Sleep(100000 * time.Hour)
fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
}