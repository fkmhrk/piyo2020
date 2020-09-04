package game

import (
	"syscall/js"
)

// Engine is an engine
type Engine interface {
	AddImage(key string, image js.Value, width, height float64)
	AddPlayerShot(shot *gameObject)
	AddEnemy(enemy *gameObject)
	AddEnemyShot(shot *gameObject)
	AddEffect(effect *gameObject)
	Player() *gameObject

	AddScore(value int)
	Miss() bool
	ToGameOver()
	DoFrame(key int16, touchDX, touchDY int, ctx js.Value)
}

const (
	gameStateMain     gameState = 1
	gameStateGameOver gameState = 2
)

type gameState int

type engine struct {
	images        map[string]*jsImage
	player        *gameObject
	playerShots   []*gameObject
	enemies       []*gameObject
	enemyShots    []*gameObject
	hiddenEnemies []*gameObject
	effects       []*gameObject
	stage         *gameObject
	gameState     gameState
	life          int
	score         int
	displayScore  int
}

// New creates engine instance
func New() Engine {
	player := newObject(objTypePlayer, 160, 440)
	player.deadFunc = deadExplode
	player.size = 4
	player.drawFunc = drawPlayer
	player.imageName = "player"
	stg := newObject(objTypeStage, 0, 0)
	stg.moveFunc = moveSequential
	stg.seqMoveFuncs = stage1Seq
	return &engine{
		images:        make(map[string]*jsImage),
		player:        player,
		enemies:       make([]*gameObject, 0, 100),
		enemyShots:    make([]*gameObject, 0, 100),
		hiddenEnemies: make([]*gameObject, 0, 100),
		effects:       make([]*gameObject, 0, 100),
		stage:         stg,
		gameState:     gameStateMain,
		life:          3,
		score:         0,
		displayScore:  0,
	}
}

func (e *engine) AddImage(key string, image js.Value, width, height float64) {
	e.images[key] = &jsImage{
		value:   image,
		width:   width,
		height:  height,
		width2:  width * 2,
		height2: height * 2,
	}
}

func (e *engine) AddPlayerShot(shot *gameObject) {
	e.playerShots = append(e.playerShots, shot)
}

func (e *engine) AddEnemy(enemy *gameObject) {
	e.enemies = append(e.enemies, enemy)
}

func (e *engine) AddEnemyShot(shot *gameObject) {
	e.enemyShots = append(e.enemyShots, shot)
}

func (e *engine) AddEffect(effect *gameObject) {
	e.effects = append(e.effects, effect)
}

func (e *engine) Player() *gameObject {
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
	e.stage.moveFunc(e.stage, e)
	e.displayScore = calcDisplayScore(e.score, e.displayScore)

	ctx.Call("clearRect", 0, 0, 320, 480)
	drawObjects(ctx, e.images, e.enemies)
	drawObjects(ctx, e.images, e.enemyShots)
	drawObjects(ctx, e.images, e.effects)
	drawObjects(ctx, e.images, e.playerShots)
	e.player.drawFunc(ctx, e.player, e.images)
	drawScore(ctx, e.images, e.displayScore, e.life)

	e.enemies = pack(e.enemies)
	e.enemyShots = pack(e.enemyShots)
	e.hiddenEnemies = pack(e.hiddenEnemies)
	e.effects = pack(e.effects)
	e.playerShots = pack(e.playerShots)
}

func movePlayer(player *gameObject, key int16, touchDX, touchDY int, engine Engine) {
	if !player.alive {
		return
	}
	if key&1 != 0 {
		player.y -= 4
	}
	if key&2 != 0 {
		player.y += 4
	}
	if key&4 != 0 {
		player.x -= 4
	}
	if key&8 != 0 {
		player.x += 4
	}
	player.x += float64(touchDX)
	player.y += float64(touchDY)
	if player.x < 0 {
		player.x = 0
	} else if player.x > 320 {
		player.x = 320
	}
	if player.y < 0 {
		player.y = 0
	} else if player.y > 480 {
		player.y = 480
	}

	if player.shotFrame > 0 {
		player.shotFrame--
	}
	if key&16 != 0 {
		if player.shotFrame == 0 {
			player.shotFrame = 5
			shot := newObject(objTypePlayerShot, player.x, player.y)
			shot.moveFunc = moveLine
			shot.vx = 0
			shot.vy = -6
			shot.size = 2
			shot.drawFunc = drawFillArc
			engine.AddPlayerShot(shot)
		}
	}
}

func moveEnemy(enemies []*gameObject, engine Engine) {
	for _, e := range enemies {
		if e.moveFunc != nil {
			e.moveFunc(e, engine)
		}
		if e.shotFunc != nil {
			e.shotFunc(e, engine)
		}
	}
}

func hitCheck(
	player *gameObject,
	playerShots []*gameObject,
	enemies []*gameObject,
	enemyShots []*gameObject,
	engine Engine,
) {
	hitCheckPlayerToEnemies(player, enemies, engine)
	hitCheckPlayerToEnemies(player, enemyShots, engine)
	hitCheckShotsToTargets(playerShots, enemies, engine)
}

func hitCheckPlayerToEnemies(player *gameObject, enemies []*gameObject, engine Engine) {
	if !player.alive {
		return
	}
	if player.frame > 0 {
		return
	}
	for _, e := range enemies {
		if isHit(player, e) {
			player.alive = false
			player.frame = 60

			player.deadFunc(engine, player)
			break
		}
	}
}

func hitCheckShotsToTargets(shots []*gameObject, targets []*gameObject, engine Engine) {
	for _, s := range shots {
		for _, t := range targets {
			if isHit(s, t) {
				s.alive = false
				t.alive = false
				engine.AddScore(t.score)
				if t.deadFunc != nil {
					t.deadFunc(engine, t)
				}
				break
			}
		}
	}
}

func checkPlayerIsDead(player *gameObject, engine Engine) {
	if player.frame > 0 {
		// blink frame
		player.frame--
	}
	if player.alive {
		return
	}
	// dead
	if player.frame == 0 {
		// game over check
		if engine.Miss() {
			player.alive = true
			player.frame = 120
			player.x = 160
			player.y = 440
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
	if score-displayScore > 1 {
		displayScore++
	}
	return displayScore
}

func drawObjects(ctx js.Value, images map[string]*jsImage, objs []*gameObject) {
	for _, o := range objs {
		o.drawFunc(ctx, o, images)
	}
}

func drawScore(ctx js.Value, images map[string]*jsImage, score int, life int) {
	x := 18 * 8
	numImage := images["number"]
	for {
		digit := score % 10
		ctx.Call("drawImage", numImage.value, digit*36, 0, 36, 36, x, 0, 18, 18)
		x -= 18
		score /= 10
		if x < 0 || score <= 0 {
			break
		}
	}
	image := images["heart"]
	for x := 0; x < life; x++ {
		ctx.Call("drawImage", image.value, 0, 0, 36, 36, x*18, 18, 18, 18)
	}
}

func pack(list []*gameObject) []*gameObject {
	pt := 0
	size := len(list)
	for i := 0; i < size; i++ {
		if list[i].alive {
			list[pt] = list[i]
			pt++
		}
	}
	return list[:pt]
}
