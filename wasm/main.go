package main

import (
	"fmt"
	"syscall/js"

	"github.com/fkmhrk/go-wasm-stg/game"
)

var (
	key         int16
	ctx         js.Value
	playerImage js.Value
)

func main() {
	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "canvas")
	ctx = canvas.Call("getContext", "2d")

	engine := game.New()

	initListeners(document, engine)

	startFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		start(engine, ctx)
		return nil
	})
	js.Global().Set("start", startFunc)

	fmt.Println("Initialization is OK")
	done := make(chan struct{}, 0)

	<-done
}

func initListeners(document js.Value, engine game.Engine) {
	setImage := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		key := args[0].JSValue().String()
		value := args[1].JSValue()

		engine.AddImage(key, value)

		return nil
	})
	js.Global().Set("setImage", setImage)

	keyDownEvent := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		keyName := event.Get("key").String()
		if keyName == "ArrowUp" {
			key |= 1
		}
		if keyName == "ArrowDown" {
			key |= 2
		}
		if keyName == "ArrowLeft" {
			key |= 4
		}
		if keyName == "ArrowRight" {
			key |= 8
		}
		if keyName == " " {
			key |= 16
		}
		return nil
	})

	keyUpEvent := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		keyName := event.Get("key").String()
		if keyName == "ArrowUp" {
			key &= ^1
		}
		if keyName == "ArrowDown" {
			key &= ^2
		}
		if keyName == "ArrowLeft" {
			key &= ^4
		}
		if keyName == "ArrowRight" {
			key &= ^8
		}
		if keyName == " " {
			key &= ^16
		}
		return nil
	})

	document.Call("addEventListener", "keydown", keyDownEvent)
	document.Call("addEventListener", "keyup", keyUpEvent)
}

func start(engine game.Engine, ctx js.Value) {
	var renderFrame js.Func
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		engine.DoFrame(key, ctx)

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})

	js.Global().Call("requestAnimationFrame", renderFrame)
}
