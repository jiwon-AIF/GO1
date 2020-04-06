package cal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	run()
}

func run() {
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "C" {
		fmt.Println("C")
		fmt.Println(Add(1, 2))
		fmt.Println(Sub(2, 1))
		fmt.Println(Mul(2, 2))
		fmt.Println(Div(4, 2))
	}
}

func Add(n1, n2 int) int {
	return n1 + n2
}}
func Sub(n1, n2 int) int {
	return n1 - n2
}
func Mul(n1, n2 int) int {
	return n1 * n2
}
func Div(n1, n2 int) int {
	return n1 / n2
}
}
