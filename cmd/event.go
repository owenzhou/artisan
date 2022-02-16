package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// makeCmd represents the make command
var makeEventCmd = &cobra.Command{
	Use:   "make:event [string to make:event]",
	Short: "Create a new event struct",
	Long:  `Create a new event struct`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("Please enter event name")
			return
		}
		fmt.Println(makeEvent(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(makeEventCmd)
}
