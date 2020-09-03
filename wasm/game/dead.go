package game

import "math"

type deadFunc func(engine Engine, obj *gameObject)

func deadExplode(engine Engine, obj *gameObject) {
	for i := 0; i < 6; i++ {
		rad := math.Pi * float64(i) / 3.0
		e1 := newObject(objTypeEnemyShot, obj.x, obj.y)
		e1.moveFunc = moveLine
		e1.vx = math.Cos(rad) * 4
		e1.vy = math.Sin(rad) * 4
		engine.AddEffect(e1)
	}
}
