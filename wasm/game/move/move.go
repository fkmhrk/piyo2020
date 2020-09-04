package move

import (
	"math"

	"github.com/fkmhrk/go-wasm-stg/game"
)

func Sequential(obj *game.GameObject, engine game.Engine) {
	obj.Frame++
	frame := obj.Frame
	for i := 0; i < len(obj.SeqMoveFuncs); i++ {
		s := obj.SeqMoveFuncs[i]
		if frame <= s.Frame {
			s.Func(obj, engine, frame)
			return
		}
		frame -= s.Frame
	}
	// reset
	obj.Frame = 0
}

func Nop(obj *game.GameObject, engine game.Engine, frame int) {
	// nop!
}

func Line(obj *game.GameObject, engine game.Engine) {
	obj.X += obj.Vx
	obj.Y += obj.Vy
	if isOutOfScreen(obj) {
		obj.Alive = false
	}
}

func LineWithFrame(obj *game.GameObject, engine game.Engine, frame int) {
	obj.X += obj.Vx
	obj.Y += obj.Vy
	if isOutOfScreen(obj) {
		obj.Alive = false
	}
}

func Sin(obj *game.GameObject, engine game.Engine, frame int) {
	rad := math.Pi * 2 * float64(frame) / 240
	obj.X += math.Sin(rad) * 2
	obj.Y += obj.Vy
	if isOutOfScreen(obj) {
		obj.Alive = false
	}
}

func Cos(obj *game.GameObject, engine game.Engine, frame int) {
	rad := math.Pi * 2 * float64(frame) / 240
	obj.X += math.Cos(rad) * 2
	obj.Y += obj.Vy
	if isOutOfScreen(obj) {
		obj.Alive = false
	}
}

func FrameUp(obj *game.GameObject, engine game.Engine) {
	obj.Frame++
	if obj.Frame > 30 {
		obj.Alive = false
	}
}

func StopAim(obj *game.GameObject, engine game.Engine) {
	obj.Frame++
	if obj.Frame < 60 {
		Line(obj, engine)
		return
	}
	if obj.Frame < 90 {
		return
	}
	if obj.Frame == 90 {
		p := engine.Player()
		rad := math.Atan2(p.Y-obj.Y, p.X-obj.X)
		obj.Vx = math.Cos(rad) * 4
		obj.Vy = math.Sin(rad) * 4
		obj.MoveFunc = Line
		return
	}
}

func isOutOfScreen(obj *game.GameObject) bool {
	if obj.X < -16 || obj.X > 336 {
		return true
	}
	if obj.Y < -16 || obj.Y > 496 {
		return true
	}
	return false
}
