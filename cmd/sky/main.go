package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(`
/ _\ | ___   _ 
\ \| |/ / | | |
_\ \   <| |_| |
\__/_|\_\\__, |
	  |___/ 
	`)

	err := rootCmd.Execute()
	if err != nil {
		log.Panicf("Error exectuing sky command: %s\n", err)
	}
}
