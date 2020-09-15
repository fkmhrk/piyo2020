package move

import (
	"math"

	"github.com/fkmhrk/go-wasm-stg/game"
)

type moveAPI struct{}

// New creates move API
func New() game.Move {
	return &moveAPI{}
}

func (m *moveAPI) Sequential() game.MoveFunc {
	return sequential
}

func (m *moveAPI) Nop() game.SeqMoveFunc {
	return nop
}

func (m *moveAPI) Line() game.MoveFunc {
	return line
}

func (m *moveAPI) FrameCountUp() game.MoveFunc {
	return frameCountUp
}

func (m *moveAPI) ItemDrop() game.MoveFunc {
	return itemDrop
}

func (m *moveAPI) LineReflect() game.MoveFunc {
	return lineReflect
}

func sequential(obj *game.GameObject, engine game.Engine) {
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

func nop(obj *game.GameObject, engine game.Engine, frame int) {
	// nop!
}

func line(obj *game.GameObject, engine game.Engine) {
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

func lineReflect(obj *game.GameObject, engine game.Engine) {
	obj.X += obj.Vx
	obj.Y += obj.Vy
	if obj.X < 0 {
		obj.Vx = -obj.Vx
		obj.X += obj.Vx
	}
	if obj.X > 320 {
		obj.Vx = -obj.Vx
		obj.X += obj.Vx
	}
	if obj.Y < 0 {
		obj.Vy = -obj.Vy
		obj.Y += obj.Vy
	}
	if isOutOfScreen(obj) {
		obj.Alive = false
	}
}

func itemDrop(obj *game.GameObject, engine game.Engine) {
	obj.X += obj.Vx
	obj.Y += obj.Vy
	obj.Vy += 0.04
	if obj.Vy > 4 {
		obj.Vy = 4
	}
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

func frameCountUp(obj *game.GameObject, engine game.Engine) {
	obj.Frame++
	if obj.Frame > 30 {
		obj.Alive = false
	}
}

func StopAim(obj *game.GameObject, engine game.Engine) {
	obj.Frame++
	if obj.Frame < 60 {
		line(obj, engine)
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
		obj.MoveFunc = line
		return
	}
}

func SlowAfter60(obj *game.GameObject, engine game.Engine) {
	obj.Frame++
	if obj.Frame < 60 {
		line(obj, engine)
		return
	}
	if obj.Frame == 60 {
		obj.Vy /= 2
		line(obj, engine)
		return
	}
	line(obj, engine)
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

var (
	SeqStage game.SeqMoveFuncs = game.SeqMoveFuncs{
		&game.SeqMove{
			Frame: 60,
			Func:  LineWithFrame,
		},
		&game.SeqMove{
			Frame: 1,
			Func: func(obj *game.GameObject, engine game.Engine, frame int) {
				obj.Vx = 0
				obj.Vy = -2
			},
		},
		&game.SeqMove{
			Frame: 9999,
			Func:  LineWithFrame,
		},
	}
)
