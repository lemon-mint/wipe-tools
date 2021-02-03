package main

import (
	"os"

	"github.com/lemon-mint/wipe-file-go/wiper"
)

func main() {
	if len(os.Args) >= 2 {
		for i := range os.Args[1:] {
			wiper.Wipe7pass(os.Args[1:][i])
		}
	}
}
