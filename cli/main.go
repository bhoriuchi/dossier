package main

import (
	"fmt"
	"os"
)

func main() {
	rootCmd.AddCommand(lookupCmd)
	initLookupCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
