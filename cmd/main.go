package main

import (
	"fmt"
	"log"

	"github.com/juancsr/go-install-with-release/cmd/cli"
)

func main() {
	fmt.Println(`
/ _\ | ___   _ 
\ \| |/ / | | |
_\ \   <| |_| |
\__/_|\_\\__, |
	  |___/ 
	`)

	err := cli.Execute()
	if err != nil {
		log.Panicf("Error exectuing sky command: %s\n", err)
	}
}
