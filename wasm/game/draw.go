package game

import (
	"math"
	"syscall/js"
)

type jsImage struct {
	value   js.Value
	width   float64
	height  float64
	width2  float64
	height2 float64
}

type drawFunc func(ctx js.Value, obj *gameObject, images map[string]*jsImage)

func drawStatic(ctx js.Value, obj *gameObject, images map[string]*jsImage) {
	image := images[obj.imageName]
	ctx.Call("drawImage", image.value, obj.x-image.width, obj.y-image.height)
}

func drawStrokeArc(ctx js.Value, obj *gameObject, images map[string]*jsImage) {
	ctx.Call("beginPath")
	ctx.Call("arc", obj.x, obj.y, 4, 0, math.Pi*2, true)
	ctx.Call("stroke")
}

func drawExpandingStrokeArc(ctx js.Value, obj *gameObject, images map[string]*jsImage) {
	ctx.Call("beginPath")
	ctx.Call("arc", obj.x, obj.y, obj.frame, 0, math.Pi*2, true)
	ctx.Call("stroke")
}

func drawFillArc(ctx js.Value, obj *gameObject, images map[string]*jsImage) {
	ctx.Call("beginPath")
	ctx.Call("arc", obj.x, obj.y, 4, 0, math.Pi*2, true)
	ctx.Call("fill")
}

func drawPlayer(ctx js.Value, obj *gameObject, images map[string]*jsImage) {
	if obj.alive {
		if obj.frame%3 == 1 {
			// blink
			return
		}
		image := images[obj.imageName]
		ctx.Call("drawImage", image.value, 0, 0, image.width2, image.height2,
			obj.x-image.width, obj.y-image.height, image.width2, image.height2)
		return
	}
}
