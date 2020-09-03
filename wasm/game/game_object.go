package game

import "math"

const (
	objTypePlayer     = 1
	objTypePlayerShot = 2
	objTypeEnemy      = 3
	objTypeEnemyShot  = 4
	objTypeStage      = 5
)

type moveFunc func(obj *gameObject, engine Engine)

type gameObject struct {
	objType  int
	x        float64
	y        float64
	vx       float64
	vy       float64
	alive    bool
	frame    int
	moveFunc moveFunc
}

func newObject(objType int, x, y float64) *gameObject {
	return &gameObject{
		objType:  objType,
		x:        x,
		y:        y,
		alive:    true,
		moveFunc: nil,
	}
}

func isHit(obj1 *gameObject, obj2 *gameObject) bool {
	xDiff := obj1.x - obj2.x
	yDiff := obj1.y - obj2.y
	dist := math.Sqrt(xDiff*xDiff + yDiff*yDiff)
	return dist < 4
}
