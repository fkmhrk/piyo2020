package game

import (
	"math"
	"syscall/js"
)

type drawFunc func(ctx js.Value, obj *gameObject, images map[string]js.Value)

func drawStatic(ctx js.Value, obj *gameObject, images map[string]js.Value) {
	ctx.Call("drawImage", images[obj.imageName], obj.x-12, obj.y-12)
}

func drawStrokeArc(ctx js.Value, obj *gameObject, images map[string]js.Value) {
	ctx.Call("beginPath")
	ctx.Call("arc", obj.x, obj.y, 4, 0, math.Pi*2, true)
	ctx.Call("stroke")
}

func drawFillArc(ctx js.Value, obj *gameObject, images map[string]js.Value) {
	ctx.Call("beginPath")
	ctx.Call("arc", obj.x, obj.y, 4, 0, math.Pi*2, true)
	ctx.Call("fill")
}

func drawPlayer(ctx js.Value, obj *gameObject, images map[string]js.Value) {
	if obj.alive {
		if obj.frame%3 == 1 {
			// blink
			return
		}
		ctx.Call("drawImage", images[obj.imageName], 0, 0, 48, 48, obj.x-24, obj.y-24, 48, 48)
		return
	}
}
