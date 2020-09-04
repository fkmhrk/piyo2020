package game

import "math/rand"

var (
	stage1Seq seqMoveFuncs = seqMoveFuncs{
		&seqMove{frame: 120, f: moveNop},
		&seqMove{frame: 30, f: stage1_1},
		&seqMove{frame: 120, f: moveNop},
		&seqMove{frame: 30, f: stage1_2},
		&seqMove{frame: 90, f: moveNop},
		&seqMove{frame: 30, f: stage1_3},
		&seqMove{frame: 120, f: moveNop},
		&seqMove{frame: 30, f: stage1_2},
		&seqMove{frame: 30, f: stage1_1},
		&seqMove{frame: 90, f: moveNop},
		&seqMove{frame: 240, f: stage1_4},
		&seqMove{frame: 60, f: moveNop},
		&seqMove{frame: 30, f: stage1_1},
		&seqMove{frame: 60, f: moveNop},
		&seqMove{frame: 30, f: stage1_2},
	}
)

func stage1_1(obj *gameObject, engine Engine, frame int) {
	if frame%5 == 0 {
		x := frame / 5
		newEnemy := newObject(objTypeEnemy, float64(16+24*x), 0)
		newEnemy.moveFunc = moveStopAim
		newEnemy.vy = 3
		newEnemy.deadFunc = deadSoloExplode
		newEnemy.score = 100
		newEnemy.size = 8
		newEnemy.drawFunc = drawStatic
		newEnemy.imageName = "player"
		newEnemy.shotFunc = shotSequential
		newEnemy.shotFrame = 0
		newEnemy.seqShotFuncs = shotSeqWaitAim
		engine.AddEnemy(newEnemy)
	}
}

func stage1_2(obj *gameObject, engine Engine, frame int) {
	if frame%5 == 0 {
		x := frame / 5
		newEnemy := newObject(objTypeEnemy, float64(304-24*x), 0)
		newEnemy.moveFunc = moveStopAim
		newEnemy.vy = 3
		newEnemy.deadFunc = deadSoloExplode
		newEnemy.score = 100
		newEnemy.size = 8
		newEnemy.drawFunc = drawStatic
		newEnemy.imageName = "player"
		newEnemy.shotFunc = shotSequential
		newEnemy.shotFrame = 0
		newEnemy.seqShotFuncs = shotSeqWaitAim
		engine.AddEnemy(newEnemy)
	}
}

func stage1_3(obj *gameObject, engine Engine, frame int) {
	if frame%5 == 0 {
		x := frame / 5
		newEnemy := newObject(objTypeEnemy, float64(304-24*x), 0)
		newEnemy.moveFunc = moveStopAim
		newEnemy.vy = 3
		newEnemy.deadFunc = deadSoloExplode
		newEnemy.score = 100
		newEnemy.size = 8
		newEnemy.drawFunc = drawStatic
		newEnemy.imageName = "player"
		newEnemy.shotFunc = shotSequential
		newEnemy.shotFrame = 0
		newEnemy.seqShotFuncs = shotSeqWaitAim
		engine.AddEnemy(newEnemy)

		newEnemy = newObject(objTypeEnemy, float64(16+24*x), 0)
		newEnemy.moveFunc = moveStopAim
		newEnemy.vy = 3
		newEnemy.deadFunc = deadSoloExplode
		newEnemy.score = 100
		newEnemy.size = 8
		newEnemy.drawFunc = drawStatic
		newEnemy.imageName = "player"
		newEnemy.shotFunc = shotSequential
		newEnemy.shotFrame = 0
		newEnemy.seqShotFuncs = shotSeqWaitAim
		engine.AddEnemy(newEnemy)
	}
}

func stage1_4(obj *gameObject, engine Engine, frame int) {
	if frame%15 == 0 {
		x := rand.Float64() * 320
		newEnemy := newObject(objTypeEnemy, x, 0)
		newEnemy.moveFunc = moveStopAim
		newEnemy.vy = 3
		newEnemy.deadFunc = deadSoloExplode
		newEnemy.score = 100
		newEnemy.size = 8
		newEnemy.drawFunc = drawStatic
		newEnemy.imageName = "player"
		newEnemy.shotFunc = shotSequential
		newEnemy.shotFrame = 0
		newEnemy.seqShotFuncs = shotSeqWaitAim
		engine.AddEnemy(newEnemy)
	}
}
