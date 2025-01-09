package main

import "testing"

func BencMarkMyEchoFunction(bf *testing.B) {
	var args []string = []string{"Something", "Is", "Amazing"}
	for i := 0; i < 1000; i++ {
		MyEcho(args)
	}
}

func BencMarkOriginalEchoFunction(bf *testing.B) {
	var args []string = []string{"Something", "Is", "Amazing"}
	for i := 0; i < 1000; i++ {
		ExampleEcho(args)
	}
}
