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
		&game.SeqMove{Frame: 1, Func: step1_1},
		&game.SeqMove{Frame: 120, Func: step1_2},
		&game.SeqMove{Frame: 120, Func: move.Nop},
		&game.SeqMove{Frame: 1, Func: step2_1},
		&game.SeqMove{Frame: 120, Func: step2_2},
	}
)

func step1_1(obj *game.GameObject, engine game.Engine, frame int) {
	newEnemy := game.NewObject(game.ObjTypeEnemy, 240, 0)
	newEnemy.MoveFunc = move.SlowAfter60
	newEnemy.HP = 12
	newEnemy.Vy = 2
	newEnemy.DeadFunc = dead.SoloExplode
	newEnemy.Score = 500
	newEnemy.Size = 8
	newEnemy.DrawFunc = draw.Static
	newEnemy.ImageName = "player"
	newEnemy.ShotFunc = shot.Sequential
	newEnemy.ShotFrame = 0
	newEnemy.SeqShotFuncs = shot.Seq3Fan
	engine.AddEnemy(newEnemy)
}

func step1_2(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%15 != 0 {
		return
	}
	for i := 0; i < 2; i++ {
		newEnemy := game.NewObject(game.ObjTypeEnemy, float64(16+i*24), 0)
		newEnemy.MoveFunc = move.StopAim
		newEnemy.Vy = 3
		newEnemy.DeadFunc = dead.SoloExplode
		newEnemy.Score = 150
		newEnemy.Size = 8
		newEnemy.DrawFunc = draw.Static
		newEnemy.ImageName = "player"
		newEnemy.ShotFunc = shot.Sequential
		newEnemy.ShotFrame = 0
		newEnemy.SeqShotFuncs = shot.SeqWaitAim
		engine.AddEnemy(newEnemy)
	}
}

func step2_1(obj *game.GameObject, engine game.Engine, frame int) {
	newEnemy := game.NewObject(game.ObjTypeEnemy, 80, 0)
	newEnemy.MoveFunc = move.SlowAfter60
	newEnemy.HP = 12
	newEnemy.Vy = 2
	newEnemy.DeadFunc = dead.SoloExplode
	newEnemy.Score = 500
	newEnemy.Size = 8
	newEnemy.DrawFunc = draw.Static
	newEnemy.ImageName = "player"
	newEnemy.ShotFunc = shot.Sequential
	newEnemy.ShotFrame = 0
	newEnemy.SeqShotFuncs = shot.Seq3Fan
	engine.AddEnemy(newEnemy)
}

func step2_2(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%15 != 0 {
		return
	}
	for i := 0; i < 2; i++ {
		newEnemy := game.NewObject(game.ObjTypeEnemy, float64(300-i*24), 0)
		newEnemy.MoveFunc = move.StopAim
		newEnemy.Vy = 3
		newEnemy.DeadFunc = dead.SoloExplode
		newEnemy.Score = 150
		newEnemy.Size = 8
		newEnemy.DrawFunc = draw.Static
		newEnemy.ImageName = "player"
		newEnemy.ShotFunc = shot.Sequential
		newEnemy.ShotFrame = 0
		newEnemy.SeqShotFuncs = shot.SeqWaitAim
		engine.AddEnemy(newEnemy)
	}
}

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
