package draw

import (
	"math"
	"syscall/js"

	"github.com/fkmhrk/go-wasm-stg/game"
)

func Static(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
	image := images[obj.ImageName]
	ctx.Call("drawImage", image.Value, obj.X-image.Width, obj.Y-image.Height)
}

func StrokeArc(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
	ctx.Call("beginPath")
	ctx.Call("arc", obj.X, obj.Y, 4, 0, math.Pi*2, true)
	ctx.Call("stroke")
}

func ExpandingStrokeArc(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
	ctx.Call("beginPath")
	ctx.Call("arc", obj.X, obj.Y, obj.Frame, 0, math.Pi*2, true)
	ctx.Call("stroke")
}

func FillArc(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
	ctx.Call("beginPath")
	ctx.Call("arc", obj.X, obj.Y, 4, 0, math.Pi*2, true)
	ctx.Call("fill")
}

func Player(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
	if obj.Alive {
		if obj.Frame%3 == 1 {
			// blink
			return
		}
		image := images[obj.ImageName]
		ctx.Call("drawImage", image.Value, 0, 0, image.Width2, image.Height2,
			obj.X-image.Width, obj.Y-image.Height, image.Width2, image.Height2)
		return
	}
}
