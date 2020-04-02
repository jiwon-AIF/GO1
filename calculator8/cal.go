package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"log"
	"fmt"
	"math/rand"
	"time"
)

var myLogger *log.Logger

func main() {
	fmt.Println("C : 계산기 실행, L : 로또 번호 생성, Q : 종료")

	reader := bufio.NewReader(os.Stdin) // os 운영체제가 갖고 있는 표준 입력을 통해 읽기용 객체 생성
	line, _ := reader.ReadString('\n') // 에러 처리 없이 한 줄 읽기
	line = strings.TrimSpace(line) // 빈칸 및 필요 없는 문자를 없애서 line에 대입
	
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
	for {
		i := "C"
if i == "C" {
	fmt.Println("숫자를 입력하세요")
	myLogger.Println("숫자를 입력하세요")

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1, _ := strconv.ParseFloat(line, 64) // 에러 처리 없이 문자를 숫자로 컨버팅

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.ParseFloat(line, 64)
	
	fmt.Printf("입력하신 숫자는 %.2f, %.2f 입니다.", n1, n2)
	myLogger.Printf("입력하신 숫자는 %.2f, %.2f 입니다.", n1, n2)

	fmt.Println("연산자를 입력하세요")
	myLogger.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100) + 1
	
	n4 := randomNum //n4라는 새로운 변수를 만들어 1~100 사이의 난수 값 생성
	n3 := float64(n4) // float타입의 n4 값을 n3에 대입

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
	} else if line == "q" {
		break
	} else {
		fmt.Println("잘못 입력하셨습니다.")
		myLogger.Println("잘못 입력하셨습니다.")
	}
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100) + 1
	
	n4 := randomNum //n4라는 새로운 변수를 만들어 1~100 사이의 난수 값 생성
	n3 := float64(n4) // float타입의 n4 값을 n3에 대입


	j := "L"
	if j == "L" { // "L" 입력 시 1~45 사이의 중복되지 않은 6개 로또 번호 생성

	for i := 0; i < 6; i++ {
	fmt.Println(n5)
	}
	fmt.Println("추첨 완료")
}
	}
}
}