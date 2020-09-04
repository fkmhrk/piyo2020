package engine

import (
	"syscall/js"

	"github.com/fkmhrk/go-wasm-stg/game"
	"github.com/fkmhrk/go-wasm-stg/game/dead"
	"github.com/fkmhrk/go-wasm-stg/game/draw"
	"github.com/fkmhrk/go-wasm-stg/game/move"
	"github.com/fkmhrk/go-wasm-stg/game/stage/stage1"
	"github.com/fkmhrk/go-wasm-stg/game/stage/stage2"
)

const (
	gameStateMain     gameState = 1
	gameStateGameOver gameState = 2
)

type gameState int

type engine struct {
	images        map[string]*game.JsImage
	player        *game.GameObject
	playerShots   []*game.GameObject
	enemies       []*game.GameObject
	enemyShots    []*game.GameObject
	hiddenEnemies []*game.GameObject
	effects       []*game.GameObject
	stage         *game.GameObject
	gameState     gameState
	life          int
	stageCount    int
	score         int
	displayScore  int
	boss          *game.GameObject
	bossMaxHP     int
}

// New creates engine instance
func New() game.Engine {
	player := game.NewObject(game.ObjTypePlayer, 160, 440)
	player.DeadFunc = dead.Explode
	player.Size = 4
	player.DrawFunc = draw.Player
	player.ImageName = "player"
	stg := game.NewObject(game.ObjTypeStage, 0, 0)
	stg.MoveFunc = move.Sequential
	stg.SeqMoveFuncs = stage1.Seq

	e := &engine{
		images:        make(map[string]*game.JsImage),
		player:        player,
		enemies:       make([]*game.GameObject, 0, 100),
		enemyShots:    make([]*game.GameObject, 0, 100),
		hiddenEnemies: make([]*game.GameObject, 0, 100),
		effects:       make([]*game.GameObject, 0, 100),
		stage:         stg,
		gameState:     gameStateMain,
		life:          2,
		stageCount:    1,
		score:         0,
		displayScore:  0,
		boss:          nil,
	}
	e.Restart()
	return e
}

func (e *engine) AddImage(key string, image js.Value, width, height float64) {
	e.images[key] = &game.JsImage{
		Value:   image,
		Width:   width,
		Height:  height,
		Width2:  width * 2,
		Height2: height * 2,
	}
}

func (e *engine) AddPlayerShot(shot *game.GameObject) {
	e.playerShots = append(e.playerShots, shot)
}

func (e *engine) AddEnemy(enemy *game.GameObject) {
	e.enemies = append(e.enemies, enemy)
}

func (e *engine) AddEnemyShot(shot *game.GameObject) {
	e.enemyShots = append(e.enemyShots, shot)
}

func (e *engine) AddEffect(effect *game.GameObject) {
	e.effects = append(e.effects, effect)
}

func (e *engine) ShowBoss(boss *game.GameObject) {
	if boss == nil {
		e.boss = nil
		return
	}
	e.AddEnemy(boss)
	e.boss = boss
	e.bossMaxHP = boss.HP
}

func (e *engine) GoToNextStage(stage int) {
	e.stageCount++
	e.stage.Frame = 0
	switch stage {
	case 1:
		e.stage.SeqMoveFuncs = stage1.Seq
	case 2:
		e.stage.SeqMoveFuncs = stage2.Seq
	}
}

func (e *engine) Player() *game.GameObject {
	return e.player
}

func (e *engine) AddScore(value int) {
	e.score += value
}

func (e *engine) Miss() bool {
	e.life--
	return e.life >= 0
}

func (e *engine) ToGameOver() {
	e.gameState = gameStateGameOver
	document := js.Global().Get("document")
	block := document.Call("getElementById", "gameover-block")
	block.Get("style").Set("display", "flex")

	js.Global().Call("setShareText", e.stageCount, e.score)
}

func (e *engine) Restart() {
	e.player.Alive = true
	e.player.Frame = 0
	e.player.X = 160
	e.player.Y = 440

	e.enemies = e.enemies[:0]
	e.enemyShots = e.enemyShots[:0]
	e.hiddenEnemies = e.hiddenEnemies[:0]
	e.effects = e.effects[:0]
	e.stage.MoveFunc = move.Sequential
	e.stage.SeqMoveFuncs = stage1.Seq
	e.stage.Frame = 0
	e.gameState = gameStateMain
	e.life = 2
	e.score = 0
	e.stageCount = 1
	e.displayScore = 0
	e.boss = nil
}

func (e *engine) DoFrame(key int16, touchDX, touchDY int, ctx js.Value) {
	if e.gameState == gameStateMain {
		e.DoMainFrame(key, touchDX, touchDY, ctx)
		return
	}
}

func (e *engine) DoMainFrame(key int16, touchDX, touchDY int, ctx js.Value) {
	movePlayer(e.player, key, touchDX, touchDY, e)
	moveEnemy(e.enemies, e)
	moveEnemy(e.hiddenEnemies, e)
	moveEnemy(e.enemyShots, e)
	moveEnemy(e.playerShots, e)
	moveEnemy(e.effects, e)
	hitCheck(e.player, e.playerShots, e.enemies, e.enemyShots, e)
	checkPlayerIsDead(e.player, e)
	e.stage.MoveFunc(e.stage, e)
	e.displayScore = calcDisplayScore(e.score, e.displayScore)

	ctx.Call("clearRect", 0, 0, 320, 480)
	drawObjects(ctx, e.images, e.enemies)
	drawObjects(ctx, e.images, e.enemyShots)
	drawObjects(ctx, e.images, e.effects)
	drawObjects(ctx, e.images, e.playerShots)
	e.player.DrawFunc(ctx, e.player, e.images)
	drawScore(ctx, e.images, e.displayScore, e.life, e.boss, e.bossMaxHP)

	e.enemies = pack(e.enemies)
	e.enemyShots = pack(e.enemyShots)
	e.hiddenEnemies = pack(e.hiddenEnemies)
	e.effects = pack(e.effects)
	e.playerShots = pack(e.playerShots)
}

func movePlayer(player *game.GameObject, key int16, touchDX, touchDY int, engine game.Engine) {
	if !player.Alive {
		return
	}
	if key&1 != 0 {
		player.Y -= 4
	}
	if key&2 != 0 {
		player.Y += 4
	}
	if key&4 != 0 {
		player.X -= 4
	}
	if key&8 != 0 {
		player.X += 4
	}
	player.X += float64(touchDX)
	player.Y += float64(touchDY)
	if player.X < 0 {
		player.X = 0
	} else if player.X > 320 {
		player.X = 320
	}
	if player.Y < 0 {
		player.Y = 0
	} else if player.Y > 480 {
		player.Y = 480
	}

	if player.ShotFrame > 0 {
		player.ShotFrame--
	}
	if key&16 != 0 {
		if player.ShotFrame == 0 {
			player.ShotFrame = 5
			shot := game.NewObject(game.ObjTypePlayerShot, player.X, player.Y)
			shot.MoveFunc = move.Line
			shot.Vx = 0
			shot.Vy = -6
			shot.Size = 2
			shot.DrawFunc = draw.FillArc
			engine.AddPlayerShot(shot)
		}
	}
}

func moveEnemy(enemies []*game.GameObject, engine game.Engine) {
	for _, e := range enemies {
		if e.MoveFunc != nil {
			e.MoveFunc(e, engine)
		}
		if e.ShotFunc != nil {
			e.ShotFunc(e, engine)
		}
	}
}

func hitCheck(
	player *game.GameObject,
	playerShots []*game.GameObject,
	enemies []*game.GameObject,
	enemyShots []*game.GameObject,
	engine game.Engine,
) {
	hitCheckPlayerToEnemies(player, enemies, engine)
	hitCheckPlayerToEnemies(player, enemyShots, engine)
	hitCheckShotsToTargets(playerShots, enemies, engine)
}

func hitCheckPlayerToEnemies(player *game.GameObject, enemies []*game.GameObject, engine game.Engine) {
	if !player.Alive {
		return
	}
	if player.Frame > 0 {
		return
	}
	for _, e := range enemies {
		if game.IsHit(player, e) {
			player.Alive = false
			player.Frame = 60

			player.DeadFunc(engine, player)
			break
		}
	}
}

func hitCheckShotsToTargets(shots []*game.GameObject, targets []*game.GameObject, engine game.Engine) {
	for _, s := range shots {
		for _, t := range targets {
			if game.IsHit(s, t) {
				s.Alive = false
				t.HP--
				if t.HP > 0 {
					break
				}
				// killed!
				t.Alive = false
				engine.AddScore(t.Score)
				if t.DeadFunc != nil {
					t.DeadFunc(engine, t)
				}
				break
			}
		}
	}
}

func checkPlayerIsDead(player *game.GameObject, engine game.Engine) {
	if player.Frame > 0 {
		// blink frame
		player.Frame--
	}
	if player.Alive {
		return
	}
	// dead
	if player.Frame == 0 {
		// game over check
		if engine.Miss() {
			player.Alive = true
			player.Frame = 120
			player.X = 160
			player.Y = 440
		} else {
			engine.ToGameOver()
		}
	}
	return

}

func calcDisplayScore(score, displayScore int) int {
	if score == displayScore {
		return score
	}
	if score-displayScore > 1000 {
		displayScore += 100
	}
	if score-displayScore > 100 {
		displayScore += 10
	}
	if score-displayScore >= 1 {
		displayScore++
	}
	return displayScore
}

func drawObjects(ctx js.Value, images map[string]*game.JsImage, objs []*game.GameObject) {
	for _, o := range objs {
		o.DrawFunc(ctx, o, images)
	}
}

func drawScore(
	ctx js.Value,
	images map[string]*game.JsImage,
	score int,
	life int,
	boss *game.GameObject,
	bossMaxHP int,
) {
	x := 18 * 8
	numImage := images["number"]
	for {
		digit := score % 10
		ctx.Call("drawImage", numImage.Value, digit*36, 0, 36, 36, x, 0, 18, 18)
		x -= 18
		score /= 10
		if x < 0 || score <= 0 {
			break
		}
	}
	image := images["heart"]
	for x := 0; x < life; x++ {
		ctx.Call("drawImage", image.Value, 0, 0, 36, 36, x*18, 18, 18, 18)
	}
	if boss != nil {
		barSize := 288 * boss.HP / bossMaxHP
		ctx.Call("fillRect", 24, 36, barSize, 9)
	}
}

func pack(list []*game.GameObject) []*game.GameObject {
	pt := 0
	size := len(list)
	for i := 0; i < size; i++ {
		if list[i].Alive {
			list[pt] = list[i]
			pt++
		}
	}
	return list[:pt]
}
