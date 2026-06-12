package main

import (
	"sync/atomic"
	"time"

	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"github.com/getlantern/systray"
)

var typing atomic.Bool

func onActivate() {
	if !typing.CompareAndSwap(false, true) {
		return
	}
	defer typing.Store(false)

	text := string(clipboard.Read(clipboard.FmtText))
	if text == "" {
		return
	}
	time.Sleep(150 * time.Millisecond)
	typeText(text)
}

func listenHotkey() {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyV)
	if err := hk.Register(); err != nil {
		return
	}
	defer hk.Unregister()

	for range hk.Keydown() {
		go onActivate()
	}
}

func onReady() {
	systray.SetTitle("Clipboard Typer")
	systray.SetTooltip("Ctrl+Shift+V para escribir el portapapeles")
	systray.SetIcon(makeIcon())

	systray.AddMenuItem("Ctrl+Shift+V — escribir portapapeles", "").Disable()
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Salir", "Cerrar Clipboard Typer")

	go listenHotkey()

	<-mQuit.ClickedCh
	systray.Quit()
}

func main() {
	clipboard.Init()
	systray.Run(onReady, func() {})
}
