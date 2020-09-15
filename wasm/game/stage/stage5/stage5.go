package stage5

import (
	"github.com/fkmhrk/go-wasm-stg/game"
	"github.com/fkmhrk/go-wasm-stg/game/draw"
	"github.com/fkmhrk/go-wasm-stg/game/move"
)

var (
	// Seq is stage 5(Ending) pattern
	Seq game.SeqMoveFuncs = game.SeqMoveFuncs{
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: len(textLine1) * 30, Func: lineText(textLine1)},
		&game.SeqMove{Frame: len(textLine2) * 30, Func: lineText(textLine2)},
		&game.SeqMove{Frame: 120, Func: nop},
		&game.SeqMove{Frame: 600, Func: end},
		&game.SeqMove{Frame: 9999, Func: nop},
	}
)

const (
	textLine1 = "THANK YOU"
	textLine2 = "FOR PLAYING"
)

func nop(obj *game.GameObject, engine game.Engine, frame int) {
}

func lineText(texts string) func(obj *game.GameObject, engine game.Engine, frame int) {
	return func(obj *game.GameObject, engine game.Engine, frame int) {
		if frame%30 != 1 {
			return
		}
		step := frame / 30
		x := step*24 + 24
		if step >= len(texts) {
			return
		}
		letter := string(texts[step])
		if letter == " " {
			return
		}
		newEnemy := game.NewObject(game.ObjTypeEnemy, float64(x), 16)
		newEnemy.MoveFunc = move.StopAim
		newEnemy.SeqMoveFuncs = nil
		newEnemy.HP = 9999
		newEnemy.Vx = 0
		newEnemy.Vy = 3
		newEnemy.DeadFunc = engine.Dead().SoloExplode()
		newEnemy.Score = 0
		newEnemy.Size = 16
		newEnemy.DrawFunc = draw.SingleText(letter)
		newEnemy.ShotFunc = nil
		newEnemy.ShotFrame = 0
		engine.AddEnemy(newEnemy)
	}
}

func end(obj *game.GameObject, engine game.Engine, frame int) {
	if frame == 1 {
		engine.AllClearBonus()
		return
	}
	if frame == 599 {
		engine.ToGameOver()
		return
	}
}
