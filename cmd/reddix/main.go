package main

import (
	"errors"
	"fmt"
	"os"

	termbox "github.com/nsf/termbox-go"
)

func die(err error) {
	termbox.Close()
	fmt.Println(err.Error())
	os.Exit(1)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			switch rval := r.(type) {
			case error:
				die(rval)
			default:
				die(errors.New(fmt.Sprint(rval)))
			}
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		die(err)
	}
}
