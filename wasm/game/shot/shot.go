package shot

import (
	"math"

	"github.com/fkmhrk/go-wasm-stg/game"
	"github.com/fkmhrk/go-wasm-stg/game/draw"
	"github.com/fkmhrk/go-wasm-stg/game/move"
)

func Sequential(obj *game.GameObject, engine game.Engine) {
	obj.ShotFrame++
	frame := obj.ShotFrame
	for i := 0; i < len(obj.SeqShotFuncs); i++ {
		s := obj.SeqShotFuncs[i]
		if frame <= s.Frame {
			s.Func(obj, engine, frame)
			return
		}
		frame -= s.Frame
	}
	// reset
	obj.ShotFrame = 0
}

func Nop(obj *game.GameObject, engine game.Engine, frame int) {
	// nop!
}

func Aim(obj *game.GameObject, engine game.Engine, frame int) {
	p := engine.Player()
	rad := math.Atan2(p.Y-obj.Y, p.X-obj.X)
	shot := game.NewObject(game.ObjTypeEnemyShot, obj.X, obj.Y)
	shot.Vx = math.Cos(rad) * 4
	shot.Vy = math.Sin(rad) * 4
	shot.MoveFunc = move.Line
	shot.DrawFunc = draw.StrokeArc
	engine.AddEnemyShot(shot)
}

func Aim5(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%30 != 0 {
		return
	}
	p := engine.Player()
	rad := math.Atan2(p.Y-obj.Y, p.X-obj.X)
	for i := 0; i < 5; i++ {
		speed := float64(i + 1)
		shot := game.NewObject(game.ObjTypeEnemyShot, obj.X, obj.Y)
		shot.Vx = math.Cos(rad) * speed
		shot.Vy = math.Sin(rad) * speed
		shot.MoveFunc = move.Line
		shot.DrawFunc = draw.StrokeArc
		engine.AddEnemyShot(shot)
	}
}

func Fan3(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%10 != 0 {
		return
	}
	loop := frame/10 - 1
	p := engine.Player()
	rad := math.Atan2(p.Y-obj.Y, p.X-obj.X) + math.Pi*2*float64(loop)/12
	speed := float64(4)
	for i := 0; i < 3; i++ {
		shot := game.NewObject(game.ObjTypeEnemyShot, obj.X, obj.Y)
		shot.Vx = math.Cos(rad) * speed
		shot.Vy = math.Sin(rad) * speed
		shot.MoveFunc = move.Line
		shot.DrawFunc = draw.StrokeArc
		engine.AddEnemyShot(shot)
		rad += math.Pi * 2 / 3
	}
}

func Circle32(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%30 != 0 {
		return
	}
	p := engine.Player()
	rad := math.Atan2(p.Y-obj.Y, p.X-obj.X)
	speed := float64(3)
	for i := 0; i < 32; i++ {
		shot := game.NewObject(game.ObjTypeEnemyShot, obj.X, obj.Y)
		shot.Vx = math.Cos(rad) * speed
		shot.Vy = math.Sin(rad) * speed
		shot.MoveFunc = move.Line
		shot.DrawFunc = draw.StrokeArc
		engine.AddEnemyShot(shot)
		rad += math.Pi * 2 / 32
	}
}

func Way3(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%60 != 15 {
		return
	}
	p := engine.Player()
	rad := math.Atan2(p.Y-obj.Y, p.X-obj.X)
	del := math.Pi / 6
	radList := []float64{rad - del, rad, rad + del}
	speed := float64(2)
	for i := 0; i < 3; i++ {
		shot := game.NewObject(game.ObjTypeEnemyShot, obj.X, obj.Y)
		shot.Vx = math.Cos(radList[i]) * speed
		shot.Vy = math.Sin(radList[i]) * speed
		shot.MoveFunc = move.Line
		shot.DrawFunc = draw.StrokeArc
		engine.AddEnemyShot(shot)
	}
}

var (
	SeqWaitAim game.SeqShotFuncs = game.SeqShotFuncs{
		&game.SeqShot{Frame: 60, Func: Nop},
		&game.SeqShot{Frame: 1, Func: Aim},
		&game.SeqShot{Frame: 9999, Func: Nop},
	}
	Seq3Fan game.SeqShotFuncs = game.SeqShotFuncs{
		&game.SeqShot{Frame: 60, Func: Nop},
		&game.SeqShot{Frame: 9999, Func: Fan3},
	}
	Seq3Way game.SeqShotFuncs = game.SeqShotFuncs{
		&game.SeqShot{Frame: 30, Func: Nop},
		&game.SeqShot{Frame: 9999, Func: Way3},
	}
)
