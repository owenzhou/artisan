package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// makeCmd represents the make command
var makeListenerCmd = &cobra.Command{
	Use:   "make:listener [string to make:listener]",
	Short: "Create a new listener struct",
	Long:  `Create a new listener struct`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("Please enter listener name")
			return
		}
		fmt.Println(makeListener(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(makeListenerCmd)
}
