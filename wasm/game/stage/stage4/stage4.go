package stage4

import (
	"math/rand"

	"github.com/fkmhrk/go-wasm-stg/game"
	"github.com/fkmhrk/go-wasm-stg/game/draw"
	"github.com/fkmhrk/go-wasm-stg/game/image"
	"github.com/fkmhrk/go-wasm-stg/game/move"
	"github.com/fkmhrk/go-wasm-stg/game/shot"
)

var (
	// Seq is stage 4 pattern
	Seq game.SeqMoveFuncs = game.SeqMoveFuncs{
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 1, Func: stageText},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 1200, Func: step1},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 1, Func: boss},
		&game.SeqMove{Frame: 9999, Func: nop},
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
	newEnemy.DrawFunc = draw.StageText(4)
	newEnemy.ShotFunc = nil
	newEnemy.ShotFrame = 0
	engine.AddEnemy(newEnemy)
}

func step1(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%30 != 0 {
		return
	}
	y := rand.Float64() * 64
	newEnemy := makeEnemy1(engine, 16, y+64)
	newEnemy.Vx = 1
	engine.AddEnemy(newEnemy)

	y = rand.Float64() * 64
	newEnemy = makeEnemy1(engine, 304, y+64)
	newEnemy.Vx = -1
	engine.AddEnemy(newEnemy)
}

func boss(obj *game.GameObject, engine game.Engine, frame int) {
	newEnemy := game.NewObject(game.ObjTypeEnemy, 160, 0)
	newEnemy.HP = 150
	newEnemy.MoveFunc = engine.Move().Sequential()
	newEnemy.Vx = 0
	newEnemy.Vy = 1
	newEnemy.SeqMoveFuncs = moveBoss
	newEnemy.DeadFunc = deadBoss
	newEnemy.Score = 30000
	newEnemy.Size = 16
	newEnemy.DrawFunc = engine.Draw().Static()
	newEnemy.ImageName = image.BossStage4
	newEnemy.ShotFunc = engine.Shot().Sequential()
	newEnemy.ShotFrame = 0
	newEnemy.SeqShotFuncs = shotBoss
	engine.ShowBoss(newEnemy)
}

func deadBoss(engine game.Engine, obj *game.GameObject) {
	engine.Dead().Explode()
	engine.ShowBoss(nil)    // clear
	engine.GoToNextStage(5) // ending
}

var (
	moveBoss game.SeqMoveFuncs = game.SeqMoveFuncs{
		&game.SeqMove{
			Frame: 60,
			Func:  move.LineWithFrame,
		},
		&game.SeqMove{
			Frame: 9999,
			Func:  randomAim,
		},
	}
	shotBoss game.SeqShotFuncs = game.SeqShotFuncs{
		&game.SeqShot{Frame: 60, Func: shot.Nop},
		&game.SeqShot{Frame: 9999, Func: bossShot},
	}
)

func randomAim(obj *game.GameObject, engine game.Engine, frame int) {
	frame %= 120
	if frame == 0 {
		nextX := rand.Float64() * 320
		nextY := rand.Float64()*64 + 32
		obj.Vx = (nextX - obj.X) / 90
		obj.Vy = (nextY - obj.Y) / 90
	}
	if frame < 90 {
		engine.Move().Line()(obj, engine)
	}
}

func bossShot(obj *game.GameObject, engine game.Engine, frame int) {
	shot.Fan5(obj, engine, frame)
	if frame%60 != 0 {
		return
	}
	shot := game.NewObject(game.ObjTypeEnemyShot, obj.X, obj.Y)
	shot.Vx = 0
	shot.Vy = 1
	shot.MoveFunc = engine.Move().Line()
	shot.DrawFunc = engine.Draw().StrokeArc()
	engine.AddEnemyShot(shot)
}

// private functions

func makeEnemy1(engine game.Engine, x, y float64) *game.GameObject {
	enemy := game.NewObject(game.ObjTypeEnemy, x, y)
	enemy.MoveFunc = engine.Move().ItemDrop()
	enemy.Vy = -2
	enemy.DeadFunc = engine.Dead().SoloExplodeWithItem(1)
	enemy.Score = 800
	enemy.Size = 8
	enemy.DrawFunc = engine.Draw().Static()
	enemy.ImageName = image.EnemyApple
	enemy.ShotFunc = engine.Shot().Sequential()
	enemy.ShotFrame = 0
	enemy.SeqShotFuncs = shot.SeqCircle16
	return enemy
}
