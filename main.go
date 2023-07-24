package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:%s <hh:mm> <text message\n>")
	}
}
