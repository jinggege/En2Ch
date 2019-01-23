package fanyi

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"os/exec"
)

func Start(exitChan chan int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("read error:", err.Error())
			exitChan <- 0
			break
		}

		//c := exec.Command("cmd.exe", "/c", "cls")
		//c.Run()

		if strings.HasPrefix(input, "--close") {
			fmt.Println("=you send close command=")
			exitChan <- 0
		}

		res := Get(input)
		res = Format(res)
		fmt.Println(res)
		fmt.Printf(Prompt)

	}
}
