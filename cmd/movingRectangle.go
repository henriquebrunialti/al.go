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
	"time"

	"github.com/spf13/cobra"

	"al.go/terminal"
	"al.go/terminal/tcell"
	"al.go/visualizer"
	"al.go/visualizer/animations/misc"
)

// movingRectangleCmd represents the movingRectangle command
var movingRectangleCmd = &cobra.Command{
	Use:   "moving-rectangle",
	Short: "A rectangle moving around the screen. Press ESC to exit, press ENTER to start/stop the animation.",
	Long:  `A rectangle moving around the screen. Press ESC to exit, press ENTER to start/stop the animation.`,
	Run: func(cmd *cobra.Command, args []string) {
		countDown()
		ctx := context.Background()
		tc := tcell.New()
		v := visualizer.New(tc)
		mv := misc.NewMovingRectangleAnimation()
		keyEvents := make(chan terminal.KeyboardEvent)
		go tc.WaitForEvent(keyEvents)
		v.Visualize(ctx, mv, keyEvents)
	},
}

func countDown() {
	fmt.Printf("\nPress Esc to stop the animation, press Enter to Pause / Start\n")
	fmt.Printf("\nStarting...\n\n")
	for i := 5; i >= 0; i-- {
		if i > 0 {
			fmt.Printf("Please wait: %v\r", i)
		} else {
			fmt.Printf("Ready............")
		}

		time.Sleep(1 * time.Second)
	}
}

func init() {
	miscCmd.AddCommand(movingRectangleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// movingRectangleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// movingRectangleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
