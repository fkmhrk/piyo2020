package game

import (
	"math"
	"syscall/js"
)

const (
	ObjTypePlayer     = 1
	ObjTypePlayerShot = 2
	ObjTypeEnemy      = 3
	ObjTypeEnemyShot  = 4
	ObjTypeStage      = 5
)

// JsImage is an image for JavaScript
type JsImage struct {
	Value   js.Value
	Width   float64
	Height  float64
	Width2  float64
	Height2 float64
}

// GameObject is a basic object
type GameObject struct {
	ObjType int
	X       float64
	Y       float64
	Vx      float64
	Vy      float64
	Size    float64
	Alive   bool
	Score   int
	HP      int

	MoveFunc     MoveFunc
	Frame        int
	SeqMoveFuncs SeqMoveFuncs

	DeadFunc  DeadFunc
	DrawFunc  DrawFunc
	ImageName string

	ShotFunc     ShotFunc
	ShotFrame    int
	SeqShotFuncs SeqShotFuncs
}

// NewObject cretes new object
func NewObject(objType int, x, y float64) *GameObject {
	return &GameObject{
		ObjType:  objType,
		X:        x,
		Y:        y,
		Alive:    true,
		MoveFunc: nil,
	}
}

func IsHit(obj1 *GameObject, obj2 *GameObject) bool {
	xDiff := obj1.X - obj2.X
	yDiff := obj1.Y - obj2.Y
	size := obj1.Size + obj2.Size
	dist := math.Sqrt(xDiff*xDiff + yDiff*yDiff)
	return dist < size
}
