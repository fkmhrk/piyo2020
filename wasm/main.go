package main

import (
	"fmt"
	"syscall/js"

	"github.com/fkmhrk/go-wasm-stg/game"
	"github.com/fkmhrk/go-wasm-stg/game/engine"
)

const (
	appVersion = "0.0.17"
)

var (
	key         int16
	ctx         js.Value
	playerImage js.Value

	touchStartX int
	touchStartY int
	touchEndX   int
	touchEndY   int
	touchDown   bool
	touchMoved  bool
)

func main() {
	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "canvas")
	ctx = canvas.Call("getContext", "2d")

	printVersion(document)

	engine := engine.New()

	initListeners(document, canvas, engine)

	startFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Call("event", "start")
		engine.AddPlayCount()
		start(engine, ctx)
		return nil
	})
	js.Global().Set("start", startFunc)

	restartFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Call("event", "restart")
		engine.AddPlayCount()
		engine.Restart()
		return nil
	})
	js.Global().Set("restart", restartFunc)

	fmt.Println("Initialization is OK")
	done := make(chan struct{}, 0)

	<-done
}

func printVersion(document js.Value) {
	pElem := document.Call("getElementById", "version")
	pElem.Set("innerText", fmt.Sprintf("Version %s", appVersion))
}

func initListeners(document js.Value, canvas js.Value, engine game.Engine) {
	setImage := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		key := args[0].JSValue().String()
		value := args[1].JSValue()
		width := args[2].JSValue().Float()
		height := args[3].JSValue().Float()

		engine.AddImage(key, value, width, height)

		return nil
	})
	js.Global().Set("setImage", setImage)

	keyDownEvent := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]

		keyName := event.Get("key").String()
		if keyName == "ArrowUp" {
			key |= 1
			event.Call("preventDefault")
		}
		if keyName == "ArrowDown" {
			key |= 2
			event.Call("preventDefault")
		}
		if keyName == "ArrowLeft" {
			key |= 4
			event.Call("preventDefault")
		}
		if keyName == "ArrowRight" {
			key |= 8
			event.Call("preventDefault")
		}
		if keyName == " " {
			key |= 16
			event.Call("preventDefault")
		}
		return nil
	})

	keyUpEvent := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		keyName := event.Get("key").String()
		if keyName == "ArrowUp" {
			key &= ^1
			event.Call("preventDefault")
		}
		if keyName == "ArrowDown" {
			key &= ^2
			event.Call("preventDefault")
		}
		if keyName == "ArrowLeft" {
			key &= ^4
			event.Call("preventDefault")
		}
		if keyName == "ArrowRight" {
			key &= ^8
			event.Call("preventDefault")
		}
		if keyName == " " {
			key &= ^16
			event.Call("preventDefault")
		}
		return nil
	})

	touchStartEvent := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		event.Call("preventDefault")
		touches := event.Get("changedTouches")
		size := touches.Get("length").Int()

		if size == 0 {
			return nil
		}
		touch := touches.Index(0)
		touchStartX = touch.Get("pageX").Int()
		touchStartY = touch.Get("pageY").Int()
		touchEndX = touchStartX
		touchEndY = touchStartY
		touchMoved = false
		touchDown = true

		return nil
	})

	touchMoveEvent := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		event.Call("preventDefault")
		touches := event.Get("changedTouches")
		size := touches.Get("length").Int()

		if size == 0 {
			return nil
		}
		touch := touches.Index(0)
		touchEndX = touch.Get("pageX").Int()
		touchEndY = touch.Get("pageY").Int()
		touchMoved = true

		return nil
	})

	touchUpEvent := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		event.Call("preventDefault")
		touches := event.Get("changedTouches")
		size := touches.Get("length").Int()

		if size == 0 {
			return nil
		}
		_ = touches
		touchMoved = false
		touchDown = false

		return nil
	})

	document.Call("addEventListener", "keydown", keyDownEvent)
	document.Call("addEventListener", "keyup", keyUpEvent)
	canvas.Call("addEventListener", "touchstart", touchStartEvent, false)
	canvas.Call("addEventListener", "touchmove", touchMoveEvent, false)
	canvas.Call("addEventListener", "touchend", touchUpEvent, false)
}

func start(engine game.Engine, ctx js.Value) {
	var renderFrame js.Func
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		touchDX := 0
		touchDY := 0
		if touchMoved {
			touchDX = touchEndX - touchStartX
			touchDY = touchEndY - touchStartY
			touchStartX = touchEndX
			touchStartY = touchEndY
			touchMoved = false
		}
		key2 := key
		if touchDown {
			// shot
			key2 |= 16
		}
		engine.DoFrame(key2, touchDX, touchDY, ctx)

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})

	js.Global().Call("requestAnimationFrame", renderFrame)
}
