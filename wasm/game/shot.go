package game

import "math"

type shotFunc func(obj *gameObject, engine Engine)
type seqShotFunc func(obj *gameObject, engine Engine, frame int)

type seqShot struct {
	frame int
	f     seqShotFunc
}

type seqShotFuncs []*seqShot

func shotSequential(obj *gameObject, engine Engine) {
	obj.shotFrame++
	frame := obj.shotFrame
	for i := 0; i < len(obj.seqShotFuncs); i++ {
		s := obj.seqShotFuncs[i]
		if frame <= s.frame {
			s.f(obj, engine, frame)
			return
		}
		frame -= s.frame
	}
	// reset
	obj.shotFrame = 0
}

func shotNop(obj *gameObject, engine Engine, frame int) {
	// nop!
}

func shotAim(obj *gameObject, engine Engine, frame int) {
	p := engine.Player()
	rad := math.Atan2(p.y-obj.y, p.x-obj.x)
	shot := newObject(objTypeEnemyShot, obj.x, obj.y)
	shot.vx = math.Cos(rad) * 4
	shot.vy = math.Sin(rad) * 4
	shot.moveFunc = moveLine
	shot.drawFunc = drawStrokeArc
	engine.AddEnemyShot(shot)
}

func shotAim5(obj *gameObject, engine Engine, frame int) {
	if frame%30 != 0 {
		return
	}
	p := engine.Player()
	rad := math.Atan2(p.y-obj.y, p.x-obj.x)
	for i := 0; i < 5; i++ {
		speed := float64(i + 1)
		shot := newObject(objTypeEnemyShot, obj.x, obj.y)
		shot.vx = math.Cos(rad) * speed
		shot.vy = math.Sin(rad) * speed
		shot.moveFunc = moveLine
		shot.drawFunc = drawStrokeArc
		engine.AddEnemyShot(shot)
	}
}

var (
	shotSeqWaitAim seqShotFuncs = seqShotFuncs{
		&seqShot{frame: 60, f: shotNop},
		&seqShot{frame: 1, f: shotAim},
		&seqShot{frame: 9999, f: shotNop},
	}
	shotSeqStage1Boss seqShotFuncs = seqShotFuncs{
		&seqShot{frame: 60, f: shotNop},
		&seqShot{frame: 9999, f: shotAim5},
	}
)
