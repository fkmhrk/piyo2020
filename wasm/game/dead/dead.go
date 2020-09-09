package dead

import (
	"fmt"
	"math"

	"github.com/fkmhrk/go-wasm-stg/game"
)

type deadAPI struct{}

func New() game.Dead {
	return &deadAPI{}
}

func (d *deadAPI) SoloExplode() game.DeadFunc {
	return soloExplode
}

func (d *deadAPI) SoloExplodeWithItem(itemID int) game.DeadFunc {
	return soloExplodeWithItem(itemID)
}

func (d *deadAPI) Explode() game.DeadFunc {
	return explode
}

func soloExplode(engine game.Engine, obj *game.GameObject) {
	e := game.NewObject(game.ObjTypeEnemyShot, obj.X, obj.Y)
	e.MoveFunc = engine.Move().FrameCountUp()
	e.DrawFunc = engine.Draw().ExpandingStrokeArc()
	e.Frame = 0
	engine.AddEffect(e)
}

func soloExplodeWithItem(itemId int) func(engine game.Engine, obj *game.GameObject) {
	imageName := fmt.Sprintf("item%d", itemId)
	return func(engine game.Engine, obj *game.GameObject) {
		soloExplode(engine, obj)

		e := game.NewObject(game.ObjTypeItem, obj.X, obj.Y)
		e.MoveFunc = engine.Move().ItemDrop()
		e.Size = 18
		e.Vy = -2
		e.DrawFunc = engine.Draw().Static()
		e.ImageName = imageName
		e.Frame = 0
		engine.AddEnemyShot(e)
	}
}

func explode(engine game.Engine, obj *game.GameObject) {
	for i := 0; i < 6; i++ {
		rad := math.Pi * float64(i) / 3.0
		e1 := game.NewObject(game.ObjTypeEnemyShot, obj.X, obj.Y)
		e1.MoveFunc = engine.Move().Line()
		e1.Vx = math.Cos(rad) * 4
		e1.Vy = math.Sin(rad) * 4
		e1.DrawFunc = engine.Draw().StrokeArc()
		engine.AddEffect(e1)
	}
}
