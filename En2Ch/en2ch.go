package main

import (
	"fmt"
	"os"

	FY "./github/com/fanyi"
)

func main() {

	fmt.Println(FY.GetDescLine())
	fmt.Printf(FY.Prompt)
	var exitChan = make(chan int)

	go FY.Start(exitChan)

	code := <-exitChan
	os.Exit(code)

	for {
	}
}
