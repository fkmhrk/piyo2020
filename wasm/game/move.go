package game

import "math"

func moveLine(obj *gameObject, engine Engine) {
	obj.x += obj.vx
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
