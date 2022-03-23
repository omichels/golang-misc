package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var fooFlag string

func main() {
	NewRootCmd().Execute()
	fmt.Printf("value of fooFlag: %s\n", fooFlag)
}

type RootCmd struct {
	cobra.Command
}

func NewRootCmd() *RootCmd {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = cobra.Command{
		Use:   "demo",
		Short: "demo",
		Long:  `a long demo text`,
	}

	rootCmd.PersistentFlags().StringVarP(&fooFlag, "foo", "f", "default", "set foo flag")
	return &RootCmd{Command: rootCmd}
}
