package stage4

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
	Seq game.SeqMoveFuncs = game.SeqMoveFuncs{
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 1, Func: stageText},
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
	newEnemy.DrawFunc = draw.StageText(3)
	newEnemy.ShotFunc = nil
	newEnemy.ShotFrame = 0
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
	engine.GoToNextStage(1) // todo make stage 3
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
		shot.DrawFunc = draw.New().StrokeArc()
		engine.AddEnemyShot(shot)
	}
}