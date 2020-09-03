package game

import "math"

type deadFunc func(engine Engine, obj *gameObject)

func deadSoloExplode(engine Engine, obj *gameObject) {
	e := newObject(objTypeEnemyShot, obj.x, obj.y)
	e.moveFunc = moveFrameUp
	e.drawFunc = drawExpandingStrokeArc
	e.frame = 0
	engine.AddEffect(e)
}

func deadExplode(engine Engine, obj *gameObject) {
	for i := 0; i < 6; i++ {
		rad := math.Pi * float64(i) / 3.0
		e1 := newObject(objTypeEnemyShot, obj.x, obj.y)
		e1.moveFunc = moveLine
		e1.vx = math.Cos(rad) * 4
		e1.vy = math.Sin(rad) * 4
		e1.drawFunc = drawStrokeArc
		engine.AddEffect(e1)
	}
}
