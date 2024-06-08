package main

import "fmt"

type Sound struct {
	isRunning bool
}

func (s *Sound) on() {
	s.isRunning = true
	fmt.Println("Turning sound on")
}

func (s *Sound) off() {
	s.isRunning = false
	fmt.Println("Turning sound off")
}
