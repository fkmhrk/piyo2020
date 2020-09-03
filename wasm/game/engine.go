package game

import (
	"math"
	"syscall/js"
)

// Engine is an engine
type Engine interface {
	AddImage(key string, image js.Value)
	AddPlayerShot(shot *gameObject)
	AddEnemy(enemy *gameObject)
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
	images        map[string]js.Value
	player        *gameObject
	playerShots   []*gameObject
	enemies       []*gameObject
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
	stg := newObject(objTypeStage, 0, 0)
	stg.moveFunc = stage1
	return &engine{
		images:        make(map[string]js.Value),
		player:        player,
		enemies:       make([]*gameObject, 0, 100),
		hiddenEnemies: make([]*gameObject, 0, 100),
		effects:       make([]*gameObject, 0, 100),
		stage:         stg,
		gameState:     gameStateMain,
		life:          3,
		score:         0,
		displayScore:  0,
	}
}

func (e *engine) AddImage(key string, image js.Value) {
	e.images[key] = image
}

func (e *engine) AddPlayerShot(shot *gameObject) {
	e.playerShots = append(e.playerShots, shot)
}

func (e *engine) AddEnemy(enemy *gameObject) {
	e.enemies = append(e.enemies, enemy)
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
	moveEnemy(e.playerShots, e)
	moveEnemy(e.effects, e)
	hitCheck(e.player, e.playerShots, e.enemies, e)
	e.stage.moveFunc(e.stage, e)
	e.displayScore = calcDisplayScore(e.score, e.displayScore)

	ctx.Call("clearRect", 0, 0, 320, 480)
	drawEnemy(ctx, e.images, e.enemies)
	drawEffects(ctx, e.images, e.effects)
	drawShots(ctx, e.images, e.playerShots)
	drawPlayer(ctx, e.images, e.player)
	drawScore(ctx, e.images, e.displayScore, e.life)

	e.enemies = pack(e.enemies)
	e.hiddenEnemies = pack(e.hiddenEnemies)
	e.effects = pack(e.effects)
	e.playerShots = pack(e.playerShots)
}

func movePlayer(player *gameObject, key int16, touchDX, touchDY int, engine Engine) {
	if !player.alive {
		player.frame--
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
	if player.frame > 0 {
		// blink frame
		player.frame--
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
			engine.AddPlayerShot(shot)
		}
	}
}

func moveEnemy(enemies []*gameObject, engine Engine) {
	for _, e := range enemies {
		if e.moveFunc != nil {
			e.moveFunc(e, engine)
		}
	}
}

func hitCheck(player *gameObject, playerShots []*gameObject, enemies []*gameObject, engine Engine) {
	hitCheckPlayerToEnemies(player, enemies, engine)
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

func drawEnemy(ctx js.Value, images map[string]js.Value, enemies []*gameObject) {
	for _, e := range enemies {
		ctx.Call("drawImage", images["player"], e.x-12, e.y-12)
	}
}

func drawEffects(ctx js.Value, images map[string]js.Value, effects []*gameObject) {
	for _, e := range effects {
		ctx.Call("beginPath")
		ctx.Call("arc", e.x, e.y, 4, 0, math.Pi*2, true)
		ctx.Call("stroke")
	}
}

func drawShots(ctx js.Value, images map[string]js.Value, shots []*gameObject) {
	for _, e := range shots {
		ctx.Call("beginPath")
		ctx.Call("arc", e.x, e.y, 4, 0, math.Pi*2, true)
		ctx.Call("fill")
	}
}

func drawPlayer(ctx js.Value, images map[string]js.Value, player *gameObject) {
	if player.alive {
		if player.frame%3 == 1 {
			// blink
			return
		}
		ctx.Call("drawImage", images["player"], 0, 0, 48, 48, player.x-24, player.y-24, 48, 48)
		return
	}
}

func drawScore(ctx js.Value, images map[string]js.Value, score int, life int) {
	ctx.Call("fillText", score, 0, 18)
	for x := 0; x < life; x++ {
		ctx.Call("drawImage", images["heart"], 0, 0, 36, 36, x*18, 18, 18, 18)
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
