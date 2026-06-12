//go:build linux

package main

import (
	"os/exec"
	"time"
)

func typeText(text string) {
	// xdotool handles Unicode and modifier keys correctly
	cmd := exec.Command("xdotool", "type", "--delay", "40", "--clearmodifiers", "--", text)
	cmd.Run()
	time.Sleep(0)
}
