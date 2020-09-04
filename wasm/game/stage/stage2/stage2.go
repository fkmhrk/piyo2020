package stage2

import (
	"math/rand"

	"github.com/fkmhrk/go-wasm-stg/game"
	"github.com/fkmhrk/go-wasm-stg/game/dead"
	"github.com/fkmhrk/go-wasm-stg/game/draw"
	"github.com/fkmhrk/go-wasm-stg/game/move"
	"github.com/fkmhrk/go-wasm-stg/game/shot"
)

var (
	Seq game.SeqMoveFuncs = game.SeqMoveFuncs{
		&game.SeqMove{Frame: 120, Func: move.Nop},
		&game.SeqMove{Frame: 6000, Func: stage1_4},
	}
)

func stage1_4(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%15 == 0 {
		x := rand.Float64() * 320
		newEnemy := game.NewObject(game.ObjTypeEnemy, x, 0)
		newEnemy.MoveFunc = move.StopAim
		newEnemy.Vy = 3
		newEnemy.DeadFunc = dead.SoloExplode
		newEnemy.Score = 100
		newEnemy.Size = 8
		newEnemy.DrawFunc = draw.Static
		newEnemy.ImageName = "player"
		newEnemy.ShotFunc = shot.Sequential
		newEnemy.ShotFrame = 0
		newEnemy.SeqShotFuncs = shot.SeqWaitAim
		engine.AddEnemy(newEnemy)
	}
}
