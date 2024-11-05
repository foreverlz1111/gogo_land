package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mycli",
		Short: "A simple CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from CLI")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
