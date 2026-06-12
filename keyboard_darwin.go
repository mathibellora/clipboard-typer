//go:build darwin

package main

/*
#cgo LDFLAGS: -framework CoreGraphics
#include <CoreGraphics/CoreGraphics.h>
#include <stdint.h>

void typeRune(uint32_t r) {
    uint16_t buf[2];
    int n = 1;
    if (r > 0xFFFF) {
        r -= 0x10000;
        buf[0] = (uint16_t)(0xD800 + (r >> 10));
        buf[1] = (uint16_t)(0xDC00 + (r & 0x3FF));
        n = 2;
    } else {
        buf[0] = (uint16_t)r;
    }
    CGEventRef down = CGEventCreateKeyboardEvent(NULL, 0, 1);
    CGEventRef up   = CGEventCreateKeyboardEvent(NULL, 0, 0);
    CGEventKeyboardSetUnicodeString(down, n, buf);
    CGEventKeyboardSetUnicodeString(up,   n, buf);
    CGEventPost(kCGSessionEventTap, down);
    CGEventPost(kCGSessionEventTap, up);
    CFRelease(down);
    CFRelease(up);
}
*/
import "C"
import "time"

func typeText(text string) {
	for _, r := range text {
		C.typeRune(C.uint32_t(r))
		time.Sleep(40 * time.Millisecond)
	}
}
