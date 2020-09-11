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
	Player() *GameObject

	// Result
	Score() int
	AddScore(value int)
	SaveResult()
	AddPlayCount()

	// Game event
	Miss() bool
	ToGameOver()
	ShowBoss(boss *GameObject)
	GoToNextStage(stage int)

	// Called from UI
	Restart()
	DoFrame(key int16, touchDX, touchDY int, ctx js.Value)

	// Functions
	Result() *Result
	Shot() Shot
	Move() Move
	Draw() Draw
	Dead() Dead
}
