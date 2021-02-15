package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bhoriuchi/dossier/cmd"
	"github.com/bhoriuchi/dossier/config"
	"github.com/spf13/cobra"
)

var (
	lookupVars   []string
	lookupConfig string
	lookupCmd    = &cobra.Command{
		Use: "lookup",
		Run: func(ccmd *cobra.Command, args []string) {
			vars := []config.Variable{}

			for _, v := range lookupVars {
				if strings.Contains(v, "=") {
					parts := strings.SplitN(v, "=", 2)
					if len(parts) == 2 {
						vars = append(vars, config.Variable{
							Name:  parts[0],
							Value: parts[1],
						})

						continue
					}
				}

				if strings.Contains(v, ":") {
					parts := strings.SplitN(v, ":", 2)
					if len(parts) == 2 {
						vars = append(vars, config.Variable{
							Name:  parts[0],
							Value: parts[1],
						})
					}
					continue
				}
			}

			err := cmd.Lookup(config.Options{
				Config:    lookupConfig,
				Variables: vars,
				Paths:     args,
			})

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}
)

func initLookupCmd() {
	lookupCmd.PersistentFlags().StringVar(&lookupConfig, "config", "", "dossier config file")
	lookupCmd.PersistentFlags().StringArrayVar(&lookupVars, "var", []string{}, "variable value")
}
