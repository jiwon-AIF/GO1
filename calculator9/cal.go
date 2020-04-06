package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var myLogger *log.Logger
var choice string

// 200402 jin : format document
// fork?!!
// abc
func main() {

	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	run()

	myLogger.Println("End of Program")
}

func run() {
	myLogger.Print("Test")

Loop:
	fmt.Println("C : 계산기 실행, L : 로또 번호 생성, Q : 종료")
	//fmt.Scanf("%v", &choice)
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "C" {
		fmt.Println("숫자를 입력하세요")
		myLogger.Println("숫자를 입력하세요")

		reader = bufio.NewReader(os.Stdin)

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		n1, _ := strconv.ParseFloat(line, 64)

		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		n2, _ := strconv.ParseFloat(line, 64)

		fmt.Printf("입력하신 숫자는 %.2f, %.2f 입니다.", n1, n2)
		myLogger.Printf("입력하신 숫자는 %.2f, %.2f 입니다.", n1, n2)

		fmt.Println("연산자를 입력하세요")
		myLogger.Println("연산자를 입력하세요")

		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)

		n4 := rand.Intn(100) + 1 // 1~100 사이의 난수 생성
		n3 := float64(n4)

		if line == "+" {
			fmt.Printf("%.2f + %.2f + %.2f = %.2f", n1, n2, n3, n1+n2+n3)
			myLogger.Printf("%.2f + %.2f +  %.2f = %.2f", n1, n2, n3, n1+n2+n3)

		} else if line == "-" {
			fmt.Printf("%.2f - %.2f - %.2f = %.2f", n1, n2, n3, n1-n2-n3)
			myLogger.Printf("%.2f - %.2f - %.2f = %.2f", n1, n2, n3, n1-n2-n3)

		} else if line == "*" {
			fmt.Printf("%.2f * %.2f * %.2f = %.2f", n1, n2, n3, n1*n2*n3)
			myLogger.Printf("%.2f * %.2f * %.2f = %.2f", n1, n2, n3, n1*n2*n3)

		} else if line == "/" {
			fmt.Printf("%.2f / %.2f / %.2f = %.2f", n1, n2, n3, n1/n2/n3)
			myLogger.Printf("%.2f / %.2f / %.2f = %.2f", n1, n2, n3, n1/n2/n3)

		} else {
			fmt.Println("잘못 입력하셨습니다.")
			myLogger.Println("잘못 입력하셨습니다.")

		}
		goto Loop

	} else if choice == "L" {
		a := make([]int, 6) // int형 배열 선언
		fmt.Println("행운의 숫자는 :")

		for i := 0; i < len(a); i++ { //숫자 6개를 뽑기 위한 for문
			a[i] = rand.Intn(45) + 1 // 1~45 숫자 중 랜덤으로 하나를 뽑아 a[0]~a[6]에 저장

			for j := 0; j < i; j++ { // 중복제거를 위한 for문 : 현재 a[]에 저장된 랜덤숫자와 이전 a[]에 들어간 숫자 비교
				if a[i] == a[j] {
					i--
				}
			}
		}
		for k := 0; k <= 5; k++ { // 채워진 배열을 출력하기 위한 for문
			fmt.Println(a[k])
			if k == 5 {
				fmt.Println("입니다.")
				goto Loop
			}
		}
	} else if choice == "Q" {
		os.Exit(1)
	} else {
		fmt.Println("안돼요")
		goto Loop
	}
}
