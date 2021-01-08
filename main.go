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
package main

import (
	"log"
	"al.go/visualizer"
	"github.com/gdamore/tcell/v2"
)

func main() {
	v, err := visualizer.New()

	if err != nil {
		log.Printf("Error on initalizing ... %v", err)
	}

	log.Printf("---visualizer properties---\n")
	w, h := v.Screen.Size()
	log.Printf("w: %v h: %v\n",w, h)
	if w/2%2 == 1 {
		w += 2
	}
	v.Rect = visualizer.NewMovingRectangle((w-1)/2, 0, 1, 1, 1)

	quit := make(chan interface{})
	
	go func() {
		for {
			ev := v.Screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					quit<--1
				}
			}
		}
	}()
	exit := false
	for !exit {
		select {
		case <-v.Ticker.C:
			 v.Screen.Clear()
			 
			 visualizer.Move(v.Rect)
			 visualizer.DrawRect(v, v.Rect)
			 v.Screen.Show()
		case <-quit:
			exit = true
			break
		}
	}
}