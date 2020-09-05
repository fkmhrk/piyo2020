package stage1

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
		&game.SeqMove{Frame: 30, Func: stage1_1},
		&game.SeqMove{Frame: 120, Func: move.Nop},
		&game.SeqMove{Frame: 30, Func: stage1_2},
		&game.SeqMove{Frame: 90, Func: move.Nop},
		&game.SeqMove{Frame: 30, Func: stage1_3},
		&game.SeqMove{Frame: 120, Func: move.Nop},
		&game.SeqMove{Frame: 30, Func: stage1_2},
		&game.SeqMove{Frame: 30, Func: stage1_1},
		&game.SeqMove{Frame: 90, Func: move.Nop},
		&game.SeqMove{Frame: 240, Func: stage1_4},
		&game.SeqMove{Frame: 60, Func: move.Nop},
		&game.SeqMove{Frame: 30, Func: stage1_1},
		&game.SeqMove{Frame: 60, Func: move.Nop},
		&game.SeqMove{Frame: 30, Func: stage1_2},
		&game.SeqMove{Frame: 120, Func: move.Nop},
		&game.SeqMove{Frame: 1, Func: stage1Boss},
		&game.SeqMove{Frame: 6000, Func: move.Nop},
	}

	Stage2Seq game.SeqMoveFuncs = game.SeqMoveFuncs{
		&game.SeqMove{Frame: 120, Func: move.Nop},
		&game.SeqMove{Frame: 6000, Func: stage1_4},
	}
)

func stage1_1(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%5 == 0 {
		x := frame / 5
		newEnemy := game.NewObject(game.ObjTypeEnemy, float64(16+28*x), 0)
		newEnemy.MoveFunc = move.StopAim
		newEnemy.Vy = 3
		newEnemy.DeadFunc = dead.SoloExplode
		newEnemy.Score = 100
		newEnemy.Size = 8
		newEnemy.DrawFunc = draw.Static
		newEnemy.ImageName = "enemy1"
		newEnemy.ShotFunc = shot.Sequential
		newEnemy.ShotFrame = 0
		newEnemy.SeqShotFuncs = shot.SeqWaitAim
		engine.AddEnemy(newEnemy)
	}
}

func stage1_2(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%5 == 0 {
		x := frame / 5
		newEnemy := game.NewObject(game.ObjTypeEnemy, float64(304-28*x), 0)
		newEnemy.MoveFunc = move.StopAim
		newEnemy.Vy = 3
		newEnemy.DeadFunc = dead.SoloExplode
		newEnemy.Score = 100
		newEnemy.Size = 8
		newEnemy.DrawFunc = draw.Static
		newEnemy.ImageName = "enemy1"
		newEnemy.ShotFunc = shot.Sequential
		newEnemy.ShotFrame = 0
		newEnemy.SeqShotFuncs = shot.SeqWaitAim
		engine.AddEnemy(newEnemy)
	}
}

func stage1_3(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%5 == 0 {
		x := frame / 5
		newEnemy := game.NewObject(game.ObjTypeEnemy, float64(304-28*x), 0)
		newEnemy.MoveFunc = move.StopAim
		newEnemy.Vy = 3
		newEnemy.DeadFunc = dead.SoloExplode
		newEnemy.Score = 100
		newEnemy.Size = 8
		newEnemy.DrawFunc = draw.Static
		newEnemy.ImageName = "enemy1"
		newEnemy.ShotFunc = shot.Sequential
		newEnemy.ShotFrame = 0
		newEnemy.SeqShotFuncs = shot.SeqWaitAim
		engine.AddEnemy(newEnemy)

		newEnemy = game.NewObject(game.ObjTypeEnemy, float64(16+28*x), 0)
		newEnemy.MoveFunc = move.StopAim
		newEnemy.Vy = 3
		newEnemy.DeadFunc = dead.SoloExplode
		newEnemy.Score = 100
		newEnemy.Size = 8
		newEnemy.DrawFunc = draw.Static
		newEnemy.ImageName = "enemy1"
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
		newEnemy.ImageName = "enemy1"
		newEnemy.ShotFunc = shot.Sequential
		newEnemy.ShotFrame = 0
		newEnemy.SeqShotFuncs = shot.SeqWaitAim
		engine.AddEnemy(newEnemy)
	}
}

func stage1Boss(obj *game.GameObject, engine game.Engine, frame int) {
	newEnemy := game.NewObject(game.ObjTypeEnemy, 160, 0)
	newEnemy.HP = 100
	newEnemy.MoveFunc = move.Sequential
	newEnemy.Vx = 0
	newEnemy.Vy = 1
	newEnemy.SeqMoveFuncs = game.SeqMoveFuncs{
		&game.SeqMove{
			Frame: 60,
			Func:  move.LineWithFrame,
		},
		&game.SeqMove{
			Frame: 1,
			Func: func(obj *game.GameObject, engine game.Engine, frame int) {
				obj.Vy = 0
			},
		},
		&game.SeqMove{
			Frame: 9999,
			Func:  move.Cos,
		},
	}
	newEnemy.DeadFunc = deadBoss
	newEnemy.Score = 10000
	newEnemy.Size = 16
	newEnemy.DrawFunc = draw.Static
	newEnemy.ImageName = "enemy1"
	newEnemy.ShotFunc = shot.Sequential
	newEnemy.ShotFrame = 0
	newEnemy.SeqShotFuncs = shotBoss
	engine.ShowBoss(newEnemy)
}

func deadBoss(engine game.Engine, obj *game.GameObject) {
	dead.Explode(engine, obj)
	engine.ShowBoss(nil) // clear
	engine.GoToNextStage(2)
}

var (
	shotBoss game.SeqShotFuncs = game.SeqShotFuncs{
		&game.SeqShot{Frame: 60, Func: shot.Nop},
		&game.SeqShot{Frame: 9999, Func: shot.Aim5},
	}
)
