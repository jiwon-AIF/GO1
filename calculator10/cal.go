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
	}
}

func Add(n1, n2 int) int {
	return n1 + n2
}