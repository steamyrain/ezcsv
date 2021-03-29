package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ezcsv",
	Short: "ezcsv convert your csv to sql with no tears",
	Long:  `No more tears ezcsv is here. Fast and reliable csv to sql translator.`,
	Run: func(cmd *cobra.Command, args []string) {
		// do something
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
