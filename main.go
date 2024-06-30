package main

import (
	"fmt"
	//	"os"

	"github.com/alecthomas/kong"
)

var cli struct {
	Debug bool `help:"Enalbe debug mode."`
}

func main() {
	ctx := kong.Parse(&cli)
	if cli.Debug {
		fmt.Printf("Context struct: %+v\n", ctx)
		fmt.Printf("CLI struct: %+v\n", cli)
	}

}
