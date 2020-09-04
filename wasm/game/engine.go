package game

import (
	"syscall/js"
)

// Engine is an engine
type Engine interface {
	AddImage(key string, image js.Value, width, height float64)
	AddPlayerShot(shot *GameObject)
	AddEnemy(enemy *GameObject)
	AddEnemyShot(shot *GameObject)
	AddEffect(effect *GameObject)
	ShowBoss(boss *GameObject)
	GoToNextStage(stage int)
	Player() *GameObject

	AddScore(value int)
	Miss() bool
	ToGameOver()
	Restart()
	DoFrame(key int16, touchDX, touchDY int, ctx js.Value)
}
