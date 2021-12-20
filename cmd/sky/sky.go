package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sky",
	Short: "It allows you to fly on your terminal",
	Long:  "It allows you to fly on your terminal",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("Fly through the sky")
}
