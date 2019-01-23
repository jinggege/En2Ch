package fanyi

import (
	"bufio"
	"fmt"
	"os"
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

		res := Get(input)
		res = Format(res)
		fmt.Println(res)
		fmt.Printf(Prompt)

	}
}
