package stage1

import (
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
		&game.SeqMove{Frame: 300, Func: step1},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 300, Func: step2},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 300, Func: step2},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 300, Func: step1},
		&game.SeqMove{Frame: 240, Func: nop},
		&game.SeqMove{Frame: 240, Func: step3},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 240, Func: step3},
		&game.SeqMove{Frame: 180, Func: nop},
		&game.SeqMove{Frame: 1, Func: boss},
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
	newEnemy.DrawFunc = draw.StageText(1)
	newEnemy.ShotFunc = nil
	newEnemy.ShotFrame = 0
	engine.AddEnemy(newEnemy)
}

func step1(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%60 != 0 {
		return
	}
	hasItem := (frame/60)%2 == 0
	newEnemy := makeEnemy1(engine, float64(288), 0, hasItem)
	engine.AddEnemy(newEnemy)
}

func step2(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%60 != 0 {
		return
	}
	hasItem := (frame/60)%2 == 0
	newEnemy := makeEnemy1(engine, float64(32), 0, hasItem)
	engine.AddEnemy(newEnemy)
}

func step3(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%30 != 0 {
		return
	}
	x := (frame/30)*32 + 32
	hasItem := (frame/30)%2 == 0
	newEnemy := makeEnemy1(engine, float64(x), 0, hasItem)
	engine.AddEnemy(newEnemy)
}

func boss(obj *game.GameObject, engine game.Engine, frame int) {
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
	newEnemy.ImageName = image.BossStage1
	newEnemy.ShotFunc = engine.Shot().Sequential()
	newEnemy.ShotFrame = 0
	newEnemy.SeqShotFuncs = shotBoss
	engine.ShowBoss(newEnemy)
}

func deadBoss(engine game.Engine, obj *game.GameObject) {
	engine.Dead().Explode()(engine, obj)
	engine.ShowBoss(nil) // clear
	engine.GoToNextStage(2)
}

// private functions

func makeEnemy1(engine game.Engine, x, y float64, hasItem bool) *game.GameObject {
	enemy := game.NewObject(game.ObjTypeEnemy, x, y)
	enemy.MoveFunc = move.StopAim
	enemy.Vy = 2
	if hasItem {
		enemy.DeadFunc = engine.Dead().SoloExplodeWithItem(1)
	} else {
		enemy.DeadFunc = engine.Dead().SoloExplode()
	}
	enemy.Score = 200
	enemy.Size = 8
	enemy.DrawFunc = engine.Draw().Static()
	enemy.ImageName = image.EnemyApple
	enemy.ShotFunc = nil
	enemy.ShotFrame = 0
	enemy.SeqShotFuncs = shot.SeqWaitAim
	return enemy
}

var (
	shotBoss game.SeqShotFuncs = game.SeqShotFuncs{
		&game.SeqShot{Frame: 60, Func: shot.Nop},
		&game.SeqShot{Frame: 9999, Func: shot.Fan3},
	}
)
