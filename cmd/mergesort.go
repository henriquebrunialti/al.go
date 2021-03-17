/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"al.go/terminal"
	"al.go/terminal/tcell"
	"al.go/visualizer"
	"al.go/visualizer/animations/algorithms/sorting"
)

// mergesortCmd represents the mergesort command
var mergesortCmd = &cobra.Command{
	Use:   "mergesort",
	Short: "Visualization of the MergeSort algorithm",
	Long: `Visualization of the MergeSort algorithm`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mergesort called")

		ctx := context.Background()
		tc := tcell.New()
		v := visualizer.New(tc)
		mv := sorting.MS()
		keyEvents := make(chan terminal.KeyboardEvent)
		go v.Visualize(ctx, mv, keyEvents)
		tc.WaitForEvent(keyEvents)
	},
}

func init() {
	sortingCmd.AddCommand(mergesortCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergesortCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mergesortCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
