package stage2

import (
	"math/rand"

	"github.com/fkmhrk/go-wasm-stg/game"
	"github.com/fkmhrk/go-wasm-stg/game/draw"
	"github.com/fkmhrk/go-wasm-stg/game/image"
	"github.com/fkmhrk/go-wasm-stg/game/move"
	"github.com/fkmhrk/go-wasm-stg/game/shot"
)

var (
	// Seq is stage1 sequence
	Seq game.SeqMoveFuncs = game.SeqMoveFuncs{
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 1, Func: stageText},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 30, Func: stage1_1},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 30, Func: stage1_2},
		&game.SeqMove{Frame: 90, Func: nop},
		&game.SeqMove{Frame: 30, Func: stage1_3},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 30, Func: stage1_2},
		&game.SeqMove{Frame: 30, Func: stage1_1},
		&game.SeqMove{Frame: 90, Func: nop},
		&game.SeqMove{Frame: 240, Func: stage1_4},
		&game.SeqMove{Frame: 60, Func: nop},
		&game.SeqMove{Frame: 30, Func: stage1_1},
		&game.SeqMove{Frame: 60, Func: nop},
		&game.SeqMove{Frame: 30, Func: stage1_2},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 1, Func: stage1Boss},
		&game.SeqMove{Frame: 6000, Func: nop},
	}
)

func nop(obj *game.GameObject, engine game.Engine, frame int) {
}

func stageText(obj *game.GameObject, engine game.Engine, frame int) {
	newEnemy := game.NewObject(game.ObjTypeEnemy, 320, 120)
	newEnemy.MoveFunc = engine.Move().Sequential()
	newEnemy.SeqMoveFuncs = move.SeqStage
	newEnemy.HP = 9999
	newEnemy.Vx = -4
	newEnemy.DeadFunc = engine.Dead().SoloExplode()
	newEnemy.Score = 0
	newEnemy.Size = 16
	newEnemy.DrawFunc = draw.StageText(2)
	newEnemy.ShotFunc = nil
	newEnemy.ShotFrame = 0
	engine.AddEnemy(newEnemy)
}

func stage1_1(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%5 != 0 {
		return
	}
	x := frame / 5
	newEnemy := makeEnemy1(engine, float64(16+28*x), 0, x%2 == 0)
	engine.AddEnemy(newEnemy)

}

func stage1_2(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%5 != 0 {
		return
	}
	x := frame / 5
	newEnemy := makeEnemy1(engine, float64(304-28*x), 0, x%2 == 0)
	engine.AddEnemy(newEnemy)
}

func stage1_3(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%5 != 0 {
		return
	}
	x := frame / 5
	newEnemy := makeEnemy1(engine, float64(304-28*x), 0, (x%2 == 0))
	engine.AddEnemy(newEnemy)

	newEnemy = makeEnemy1(engine, float64(16+28*x), 0, (x%2 == 1))
	engine.AddEnemy(newEnemy)

}

func stage1_4(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%15 != 0 {
		return
	}
	x := rand.Float64() * 320
	newEnemy := makeEnemy1(engine, x, 0, true)
	engine.AddEnemy(newEnemy)

}

func stage1Boss(obj *game.GameObject, engine game.Engine, frame int) {
	newEnemy := game.NewObject(game.ObjTypeEnemy, 160, 0)
	newEnemy.HP = 100
	newEnemy.MoveFunc = engine.Move().Sequential()
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
	newEnemy.Size = 24
	newEnemy.DrawFunc = engine.Draw().Static()
	newEnemy.ImageName = image.BossStage2
	newEnemy.ShotFunc = engine.Shot().Sequential()
	newEnemy.ShotFrame = 0
	newEnemy.SeqShotFuncs = shotBoss
	engine.ShowBoss(newEnemy)
}

func deadBoss(engine game.Engine, obj *game.GameObject) {
	engine.Dead().Explode()(engine, obj)
	engine.ShowBoss(nil) // clear
	engine.GoToNextStage(3)
}

// private functions

func makeEnemy1(engine game.Engine, x, y float64, hasItem bool) *game.GameObject {
	enemy := game.NewObject(game.ObjTypeEnemy, x, y)
	enemy.MoveFunc = move.StopAim
	enemy.Vy = 3
	if hasItem {
		enemy.DeadFunc = engine.Dead().SoloExplodeWithItem(1)
	} else {
		enemy.DeadFunc = engine.Dead().SoloExplode()
	}
	enemy.Score = 300
	enemy.Size = 8
	enemy.DrawFunc = engine.Draw().Static()
	enemy.ImageName = image.EnemyApple
	enemy.ShotFunc = engine.Shot().Sequential()
	enemy.ShotFrame = 0
	enemy.SeqShotFuncs = shot.SeqWaitAim
	return enemy
}

var (
	shotBoss game.SeqShotFuncs = game.SeqShotFuncs{
		&game.SeqShot{Frame: 60, Func: shot.Nop},
		&game.SeqShot{Frame: 9999, Func: shot.Aim5},
	}
)
