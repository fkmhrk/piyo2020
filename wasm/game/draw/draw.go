package draw

import (
	"fmt"
	"math"
	"syscall/js"

	"github.com/fkmhrk/go-wasm-stg/game"
)

type drawAPI struct{}

func New() game.Draw {
	return &drawAPI{}
}

func (d *drawAPI) Static() game.DrawFunc {
	return static
}

func (d *drawAPI) StrokeArc() game.DrawFunc {
	return strokeArc
}

func (d *drawAPI) ExpandingStrokeArc() game.DrawFunc {
	return expandingStrokeArc
}

func (d *drawAPI) FillArc() game.DrawFunc {
	return FillArc
}

func (d *drawAPI) StageText(stage int) game.DrawFunc {
	return StageText(stage)
}

func (d *drawAPI) Player() game.DrawFunc {
	return Player
}

func static(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
	image := images[obj.ImageName]
	ctx.Call("drawImage", image.Value, obj.X-image.Width, obj.Y-image.Height)
}

func strokeArc(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
	ctx.Call("beginPath")
	ctx.Call("arc", obj.X, obj.Y, 4, 0, math.Pi*2, true)
	ctx.Call("stroke")
}

func expandingStrokeArc(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
	ctx.Call("beginPath")
	ctx.Call("arc", obj.X, obj.Y, obj.Frame, 0, math.Pi*2, true)
	ctx.Call("stroke")
}

func FillArc(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
	ctx.Call("beginPath")
	ctx.Call("arc", obj.X, obj.Y, 4, 0, math.Pi*2, true)
	ctx.Call("fill")
}

func StageText(stage int) func(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
	text := fmt.Sprintf("Stage %d", stage)
	return func(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
		ctx.Set("font", "32px sans-serif")
		ctx.Call("fillText", text, obj.X, obj.Y)
	}
}

func SingleText(text string) func(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
	return func(ctx js.Value, obj *game.GameObject, images map[string]*game.JsImage) {
		ctx.Set("font", "32px sans-serif")
		ctx.Call("fillText", text, obj.X, obj.Y)
	}
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
