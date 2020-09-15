package stage3

import (
	"math"
	"math/rand"

	"github.com/fkmhrk/go-wasm-stg/game"
	"github.com/fkmhrk/go-wasm-stg/game/draw"
	"github.com/fkmhrk/go-wasm-stg/game/image"
	"github.com/fkmhrk/go-wasm-stg/game/move"
	"github.com/fkmhrk/go-wasm-stg/game/shot"
)

var (
	// Seq is stage 3 pattern
	Seq game.SeqMoveFuncs = game.SeqMoveFuncs{
		&game.SeqMove{Frame: 180, Func: nop},
		&game.SeqMove{Frame: 1, Func: stageText},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 1, Func: step1_1},
		&game.SeqMove{Frame: 120, Func: step1_2},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 1, Func: step2_1},
		&game.SeqMove{Frame: 120, Func: step2_2},
		&game.SeqMove{Frame: 240, Func: nop},
		&game.SeqMove{Frame: 1, Func: step3_1},
		&game.SeqMove{Frame: 120, Func: step3_2},
		&game.SeqMove{Frame: 180, Func: nop},
		&game.SeqMove{Frame: 360, Func: step4},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 1, Func: step2_1},
		&game.SeqMove{Frame: 120, Func: step2_2},
		&game.SeqMove{Frame: 1, Func: step1_1},
		&game.SeqMove{Frame: 120, Func: step1_2},
		&game.SeqMove{Frame: 180, Func: nop},
		&game.SeqMove{Frame: 1, Func: step1_1},
		&game.SeqMove{Frame: 1, Func: step2_1},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 1, Func: step3_1},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 1, Func: step1_1},
		&game.SeqMove{Frame: 1, Func: step2_1},
		&game.SeqMove{Frame: 300, Func: nop},
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
	newEnemy.DrawFunc = draw.StageText(3)
	newEnemy.ShotFunc = nil
	newEnemy.ShotFrame = 0
	engine.AddEnemy(newEnemy)
}

func step1_1(obj *game.GameObject, engine game.Engine, frame int) {
	newEnemy := makeEnemy2(engine, 240, 0, true)
	engine.AddEnemy(newEnemy)
}

func step1_2(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%15 != 0 {
		return
	}
	hasItem := (frame/15)%2 == 0
	newEnemy := makeEnemy1(engine, 16, 0, hasItem)
	engine.AddEnemy(newEnemy)

	newEnemy = makeEnemy1(engine, 40, 0, hasItem)
	engine.AddEnemy(newEnemy)
}

func step2_1(obj *game.GameObject, engine game.Engine, frame int) {
	newEnemy := makeEnemy2(engine, 80, 0, true)
	engine.AddEnemy(newEnemy)
}

func step2_2(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%15 != 0 {
		return
	}
	hasItem := (frame/15)%2 == 0
	newEnemy := makeEnemy1(engine, 300, 0, hasItem)
	engine.AddEnemy(newEnemy)

	newEnemy = makeEnemy1(engine, 276, 0, hasItem)
	engine.AddEnemy(newEnemy)
}
func step3_1(obj *game.GameObject, engine game.Engine, frame int) {
	newEnemy := makeEnemy2(engine, 160, 0, true)
	engine.AddEnemy(newEnemy)
}

func step3_2(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%15 != 0 {
		return
	}
	hasItem := (frame/15)%2 == 0
	newEnemy := makeEnemy1(engine, 16, 0, hasItem)
	engine.AddEnemy(newEnemy)

	newEnemy = makeEnemy1(engine, 304, 0, hasItem)
	engine.AddEnemy(newEnemy)
}

func step4(obj *game.GameObject, engine game.Engine, frame int) {
	if frame%10 != 0 {
		return
	}
	x := rand.Float64() * 320
	vy := rand.Float64()*3 + 1

	newEnemy := makeEnemy1(engine, x, 0, true)
	newEnemy.MoveFunc = engine.Move().Line()
	newEnemy.Vy = vy
	newEnemy.Score = 150
	newEnemy.ShotFunc = engine.Shot().Sequential()
	newEnemy.ShotFrame = 0
	newEnemy.SeqShotFuncs = shot.Seq3Way
	engine.AddEnemy(newEnemy)

}

func boss(obj *game.GameObject, engine game.Engine, frame int) {
	newEnemy := game.NewObject(game.ObjTypeEnemy, 160, 0)
	newEnemy.HP = 100
	newEnemy.MoveFunc = engine.Move().Sequential()
	newEnemy.Vx = 0
	newEnemy.Vy = 1
	newEnemy.SeqMoveFuncs = moveBoss
	newEnemy.DeadFunc = deadBoss
	newEnemy.Score = 50000
	newEnemy.Size = 24
	newEnemy.DrawFunc = engine.Draw().Static()
	newEnemy.ImageName = image.BossStage3
	newEnemy.ShotFunc = engine.Shot().Sequential()
	newEnemy.ShotFrame = 0
	newEnemy.SeqShotFuncs = shotBoss
	engine.ShowBoss(newEnemy)
}

func deadBoss(engine game.Engine, obj *game.GameObject) {
	engine.Dead().Explode()
	engine.ShowBoss(nil) // clear
	engine.GoToNextStage(4)
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
	shot.Circle32(obj, engine, frame)
	if frame%150 != 50 {
		return
	}
	p := engine.Player()
	rad := math.Atan2(p.Y-obj.Y, p.X-obj.X)
	delta := math.Pi * 2 / 64
	radList := make([]float64, 17, 17)
	for i := -8; i <= 8; i++ {
		radList[i+8] = rad + delta*float64(i)
	}
	speed := float64(1)
	for i := 0; i < 17; i++ {
		shot := game.NewObject(game.ObjTypeEnemyShot, obj.X, obj.Y)
		shot.Vx = math.Cos(radList[i]) * speed
		shot.Vy = math.Sin(radList[i]) * speed
		shot.MoveFunc = engine.Move().Line()
		shot.DrawFunc = engine.Draw().StrokeArc()
		engine.AddEnemyShot(shot)
	}
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
	enemy.Score = 500
	enemy.Size = 8
	enemy.DrawFunc = engine.Draw().Static()
	enemy.ImageName = image.EnemyApple
	enemy.ShotFunc = engine.Shot().Sequential()
	enemy.ShotFrame = 0
	enemy.SeqShotFuncs = shot.SeqWaitAim
	return enemy
}

func makeEnemy2(engine game.Engine, x, y float64, hasItem bool) *game.GameObject {
	enemy := game.NewObject(game.ObjTypeEnemy, x, 0)
	enemy.MoveFunc = move.SlowAfter60
	enemy.HP = 12
	enemy.Vy = 2
	if hasItem {
		enemy.DeadFunc = engine.Dead().SoloExplodeWithItem3(1)
	} else {
		enemy.DeadFunc = engine.Dead().SoloExplode()
	}
	enemy.Score = 7500
	enemy.Size = 12
	enemy.DrawFunc = engine.Draw().Static()
	enemy.ImageName = image.EnemyBanana
	enemy.ShotFunc = engine.Shot().Sequential()
	enemy.ShotFrame = 0
	enemy.SeqShotFuncs = shot.Seq3Fan
	return enemy
}
