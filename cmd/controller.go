package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Resource bool

// makeCmd represents the make command
var makeControllerCmd = &cobra.Command{
	Use:   "make:controller [string to make:controller]",
	Short: "Create a new controller struct",
	Long:  `Create a new controller struct`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("Please enter controller name.")
			return
		}
		fmt.Println(makeController(args[0], Resource))
	},
}

func init() {
	makeControllerCmd.Flags().BoolVarP(&Resource, "resource", "r", false, "Create restfull api controller")
	rootCmd.AddCommand(makeControllerCmd)
}
