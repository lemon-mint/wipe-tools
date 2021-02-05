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

func wipe21(filename string) error {
	err := wiper.RandFill(filename, 10)
	if err != nil {
		return err
	}
	err = wiper.ZeroFill(filename, 2)
	if err != nil {
		return err
	}
	err = wiper.ComplementFill(filename)
	if err != nil {
		return err
	}
	err = wiper.RandFill(filename, 3)
	if err != nil {
		return err
	}
	err = wiper.ComplementFill(filename)
	if err != nil {
		return err
	}
	err = wiper.RandFill(filename, 3)
	if err != nil {
		return err
	}
	filename, err = wiper.MixFileName(filename, 10)
	if err != nil {
		return err
	}
	err = wiper.MixTime(filename, 10)
	if err != nil {
		return err
	}
	filename, err = wiper.MixFileName(filename, 5)
	if err != nil {
		return err
	}
	err = wiper.MixTime(filename, 5)
	if err != nil {
		return err
	}
	err = wiper.ZeroFill(filename, 1)
	if err != nil {
		return err
	}
	return os.Remove(filename)
}
