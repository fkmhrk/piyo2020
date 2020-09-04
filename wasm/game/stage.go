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
		&seqMove{frame: 120, f: moveNop},
		&seqMove{frame: 1, f: stage1Boss},
		&seqMove{frame: 6000, f: moveNop},
	}

	stage2Seq seqMoveFuncs = seqMoveFuncs{
		&seqMove{frame: 120, f: moveNop},
		&seqMove{frame: 6000, f: stage1_4},
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

func stage1Boss(obj *gameObject, engine Engine, frame int) {
	newEnemy := newObject(objTypeEnemy, 160, 0)
	newEnemy.hp = 100
	newEnemy.moveFunc = moveSequential
	newEnemy.vx = 0
	newEnemy.vy = 1
	newEnemy.seqMoveFuncs = seqMoveFuncs{
		&seqMove{
			frame: 60,
			f:     moveLineWithFrame,
		},
		&seqMove{
			frame: 1,
			f: func(obj *gameObject, engine Engine, frame int) {
				obj.vy = 0
			},
		},
		&seqMove{
			frame: 9999,
			f:     moveCos,
		},
	}
	newEnemy.deadFunc = deadStage1Boss
	newEnemy.score = 10000
	newEnemy.size = 16
	newEnemy.drawFunc = drawStatic
	newEnemy.imageName = "player"
	newEnemy.shotFunc = shotSequential
	newEnemy.shotFrame = 0
	newEnemy.seqShotFuncs = shotSeqStage1Boss
	engine.ShowBoss(newEnemy)
}
