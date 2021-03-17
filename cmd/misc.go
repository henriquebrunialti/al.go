package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

)

// miscCmd represents the misc command
var miscCmd = &cobra.Command{
	Use:   "misc",
	Short: "Miscellaneous Animations Loops",
	Long:  `Some miscellaneous annimation loops made for fun`,
	Run: func(cmd *cobra.Command, args []string) {


		fmt.Printf("These are some miscellaneous animations.")
		fmt.Printf("\nUse -h to list all supported animations.\n")

		fmt.Printf("\nANIMATIONS:\n")
		fmt.Printf("\nal.go misc moving-rectangle	 A rectangle moving around the screen")
	},
}

func init() {
	rootCmd.AddCommand(miscCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// miscCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// miscCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
