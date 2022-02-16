package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// makeCmd represents the make command
var makeModelCmd = &cobra.Command{
	Use:   "make:model [string to make:model]",
	Short: "Create a new model struct",
	Long:  `Create a new model struct`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("Please enter model name")
			return
		}
		fmt.Println(makeModel(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(makeModelCmd)
}
