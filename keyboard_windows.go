//go:build windows

package main

import (
	"syscall"
	"time"
	"unsafe"
)

var (
	user32    = syscall.NewLazyDLL("user32.dll")
	sendInput = user32.NewProc("SendInput")
)

const (
	inputKeyboard    = 1
	keyeventfUnicode = 0x0004
	keyeventfKeyUp   = 0x0002
)

// INPUT struct layout for 64-bit Windows (40 bytes total).
// The union is sized to MOUSEINPUT (32 bytes); KEYBDINPUT fits within it.
type kbInput struct {
	typ   uint32
	_     [4]byte  // align union to 8 bytes
	vk    uint16
	scan  uint16
	flags uint32
	time_ uint32
	_     [4]byte  // align dwExtraInfo to 8 bytes
	extra uintptr  // 8 bytes
	_     [8]byte  // pad union to 32 bytes (MOUSEINPUT size)
}

func sendKey(scan uint16, flags uint32) {
	in := kbInput{typ: inputKeyboard, scan: scan, flags: flags}
	sendInput.Call(1, uintptr(unsafe.Pointer(&in)), unsafe.Sizeof(in))
}

func typeText(text string) {
	for _, r := range text {
		if r > 0xFFFF {
			// Surrogate pair for characters outside BMP
			r -= 0x10000
			high := uint16(0xD800 + (r >> 10))
			low := uint16(0xDC00 + (r & 0x3FF))
			sendKey(high, keyeventfUnicode)
			sendKey(high, keyeventfUnicode|keyeventfKeyUp)
			sendKey(low, keyeventfUnicode)
			sendKey(low, keyeventfUnicode|keyeventfKeyUp)
		} else {
			sendKey(uint16(r), keyeventfUnicode)
			sendKey(uint16(r), keyeventfUnicode|keyeventfKeyUp)
		}
		time.Sleep(40 * time.Millisecond)
	}
}
