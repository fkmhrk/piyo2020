package dead

import (
	"math"

	"github.com/fkmhrk/go-wasm-stg/game"
	"github.com/fkmhrk/go-wasm-stg/game/draw"
	"github.com/fkmhrk/go-wasm-stg/game/move"
)

func SoloExplode(engine game.Engine, obj *game.GameObject) {
	e := game.NewObject(game.ObjTypeEnemyShot, obj.X, obj.Y)
	e.MoveFunc = move.FrameUp
	e.DrawFunc = draw.ExpandingStrokeArc
	e.Frame = 0
	engine.AddEffect(e)
}

func Explode(engine game.Engine, obj *game.GameObject) {
	for i := 0; i < 6; i++ {
		rad := math.Pi * float64(i) / 3.0
		e1 := game.NewObject(game.ObjTypeEnemyShot, obj.X, obj.Y)
		e1.MoveFunc = move.Line
		e1.Vx = math.Cos(rad) * 4
		e1.Vy = math.Sin(rad) * 4
		e1.DrawFunc = draw.StrokeArc
		engine.AddEffect(e1)
	}
}

func Stage1Boss(engine game.Engine, obj *game.GameObject) {
	Explode(engine, obj)
	engine.ShowBoss(nil) // clear
	engine.GoToNextStage(2)
}
