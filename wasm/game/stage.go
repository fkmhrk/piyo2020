package game

func stage1(obj *gameObject, engine Engine) {
	obj.frame++
	frame := obj.frame
	if frame < 120 {
		return
	}
	frame -= 120
	if frame < 30 {
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
			newEnemy.shotFunc = shotAim
			newEnemy.shotFrame = 0
			engine.AddEnemy(newEnemy)
		}
		return
	}
	frame -= 30
	if frame < 120 {
		return
	}
	frame -= 120
	if frame < 30 {
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
			newEnemy.shotFunc = shotAim
			newEnemy.shotFrame = 0
			engine.AddEnemy(newEnemy)
		}
		return
	}
	frame -= 30
	if frame < 120 {
		return
	}
	frame -= 120
	if frame < 30 {
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
			newEnemy.shotFunc = shotAim
			newEnemy.shotFrame = 0
			engine.AddEnemy(newEnemy)

			newEnemy = newObject(objTypeEnemy, float64(16+24*x), 0)
			newEnemy.moveFunc = moveStopAim
			newEnemy.vy = 3
			newEnemy.deadFunc = deadSoloExplode
			newEnemy.score = 100
			newEnemy.size = 8
			newEnemy.drawFunc = drawStatic
			newEnemy.imageName = "player"
			newEnemy.shotFunc = shotAim
			newEnemy.shotFrame = 0
			engine.AddEnemy(newEnemy)
		}
		return
	}
	obj.frame = 0
}
