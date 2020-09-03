package game

import "math"

type shotFunc func(obj *gameObject, engine Engine)

func shotAim(obj *gameObject, engine Engine) {
	obj.shotFrame++
	if obj.shotFrame < 60 {
		return
	}
	if obj.frame == 60 {
		p := engine.Player()
		rad := math.Atan2(p.y-obj.y, p.x-obj.x)
		shot := newObject(objTypeEnemyShot, obj.x, obj.y)
		shot.vx = math.Cos(rad) * 4
		shot.vy = math.Sin(rad) * 4
		shot.moveFunc = moveLine
		shot.drawFunc = drawStrokeArc
		engine.AddEnemyShot(shot)
		return
	}
}
