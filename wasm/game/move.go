package game

import "math"

type moveFunc func(obj *gameObject, engine Engine)
type seqMoveFunc func(obj *gameObject, engine Engine, frame int)

type seqMove struct {
	frame int
	f     seqMoveFunc
}

type seqMoveFuncs []*seqMove

func moveSequential(obj *gameObject, engine Engine) {
	obj.frame++
	frame := obj.frame
	for i := 0; i < len(obj.seqMoveFuncs); i++ {
		s := obj.seqMoveFuncs[i]
		if frame <= s.frame {
			s.f(obj, engine, frame)
			return
		}
		frame -= s.frame
	}
	// reset
	obj.frame = 0
}

func moveNop(obj *gameObject, engine Engine, frame int) {
	// nop!
}

func moveLine(obj *gameObject, engine Engine) {
	obj.x += obj.vx
	obj.y += obj.vy
	if isOutOfScreen(obj) {
		obj.alive = false
	}
}

func moveLineWithFrame(obj *gameObject, engine Engine, frame int) {
	obj.x += obj.vx
	obj.y += obj.vy
	if isOutOfScreen(obj) {
		obj.alive = false
	}
}

func moveSin(obj *gameObject, engine Engine, frame int) {
	rad := math.Pi * 2 * float64(frame) / 240
	obj.x += math.Sin(rad) * 2
	obj.y += obj.vy
	if isOutOfScreen(obj) {
		obj.alive = false
	}
}

func moveCos(obj *gameObject, engine Engine, frame int) {
	rad := math.Pi * 2 * float64(frame) / 240
	obj.x += math.Cos(rad) * 2
	obj.y += obj.vy
	if isOutOfScreen(obj) {
		obj.alive = false
	}
}

func moveFrameUp(obj *gameObject, engine Engine) {
	obj.frame++
	if obj.frame > 30 {
		obj.alive = false
	}
}

func moveStopAim(obj *gameObject, engine Engine) {
	obj.frame++
	if obj.frame < 60 {
		moveLine(obj, engine)
		return
	}
	if obj.frame < 90 {
		return
	}
	if obj.frame == 90 {
		p := engine.Player()
		rad := math.Atan2(p.y-obj.y, p.x-obj.x)
		obj.vx = math.Cos(rad) * 4
		obj.vy = math.Sin(rad) * 4
		obj.moveFunc = moveLine
		return
	}
}

func isOutOfScreen(obj *gameObject) bool {
	if obj.x < -16 || obj.x > 336 {
		return true
	}
	if obj.y < -16 || obj.y > 496 {
		return true
	}
	return false
}
