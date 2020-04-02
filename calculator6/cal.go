package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"log"
	"fmt"
)

var myLogger *log.Logger

func main() {
	
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
	panic(err)
}
defer fpLog.Close()

myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

//....
run()

myLogger.Println("End of Program")
}

func run() {
myLogger.Print("Test")

	for {
	
	fmt.Println("숫자를 입력하세요")
	myLogger.Println("숫자를 입력하세요")

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1, _ := strconv.ParseFloat(line, 64)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.ParseFloat(line, 64)

	fmt.Printf("입력하신 숫자는 %f, %f 입니다.", n1, n2)
	myLogger.Printf("입력하신 숫자는 %f, %f 입니다.", n1, n2)

	fmt.Println("연산자를 입력하세요")
	myLogger.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+" {
		fmt.Printf("%.2f + %.2f = %.2f", n1, n2, n1+n2)
		myLogger.Printf("%.2f + %.2f = %.2f", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%.2f - %.2f = %.2f", n1, n2, n1-n2)
		myLogger.Printf("%.2f - %.2f = %.2f", n1, n2, n1-n2)
	} else if line == "*" {
		fmt.Printf("%.2f * %.2f = %.2f", n1, n2, n1*n2)
		myLogger.Printf("%.2f * %.2f = %.2f", n1, n2, n1*n2)
	} else if line == "/" {
		fmt.Printf("%.2f / %.2f = %.2f", n1, n2, n1/n2)
		myLogger.Printf("%.2f / %.2f = %.2f", n1, n2, n1/n2)
	} else if line == "q" {
		break
	} else {
		fmt.Println("잘못 입력하셨습니다.")
		myLogger.Println("잘못 입력하셨습니다.")
	}
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
}
}