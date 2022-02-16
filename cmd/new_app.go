package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// makeCmd represents the make command
var newAppCmd = &cobra.Command{
	Use:   "new [string to new]",
	Short: "Create a new application",
	Long:  `Create a new application`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("Please enter app name.")
			return
		}
		fmt.Println(newApp(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(newAppCmd)
}
