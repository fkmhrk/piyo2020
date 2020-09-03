package game

import (
	"syscall/js"
)

// Engine is an engine
type Engine interface {
	AddImage(key string, image js.Value)
	AddEnemy(enemy *gameObject)
	Player() *gameObject

	Miss() bool
	ToGameOver()
	DoFrame(key int16, ctx js.Value)
}

const (
	gameStateMain     gameState = 1
	gameStateGameOver gameState = 2
)

type gameState int

type engine struct {
	images        map[string]js.Value
	player        *gameObject
	enemies       []*gameObject
	hiddenEnemies []*gameObject
	stage         *gameObject
	gameState     gameState
	life          int
}

// New creates engine instance
func New() Engine {
	stg := newObject(objTypeStage, 0, 0)
	stg.moveFunc = stage1
	return &engine{
		images:        make(map[string]js.Value),
		player:        newObject(objTypePlayer, 160, 440),
		enemies:       make([]*gameObject, 0, 100),
		hiddenEnemies: make([]*gameObject, 0, 100),
		stage:         stg,
		gameState:     gameStateMain,
		life:          3,
	}
}

func (e *engine) AddImage(key string, image js.Value) {
	e.images[key] = image
}

func (e *engine) AddEnemy(enemy *gameObject) {
	e.enemies = append(e.enemies, enemy)
}

func (e *engine) Player() *gameObject {
	return e.player
}

func (e *engine) Miss() bool {
	e.life--
	return e.life > 0
}

func (e *engine) ToGameOver() {
	e.gameState = gameStateGameOver
}

func (e *engine) DoFrame(key int16, ctx js.Value) {
	if e.gameState == gameStateMain {
		e.DoMainFrame(key, ctx)
		return
	}
}

func (e *engine) DoMainFrame(key int16, ctx js.Value) {
	movePlayer(e.player, key, e)
	moveEnemy(e.enemies, e)
	moveEnemy(e.hiddenEnemies, e)
	hitCheck(e.player, e.enemies)
	e.stage.moveFunc(e.stage, e)

	ctx.Call("clearRect", 0, 0, 320, 480)
	drawEnemy(ctx, e.images, e.enemies)
	drawPlayer(ctx, e.images, e.player)

	e.enemies = pack(e.enemies)
	e.hiddenEnemies = pack(e.hiddenEnemies)
}

func movePlayer(player *gameObject, key int16, engine Engine) {
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
}

func moveEnemy(enemies []*gameObject, engine Engine) {
	for _, e := range enemies {
		if e.moveFunc != nil {
			e.moveFunc(e, engine)
		}
	}
}

func hitCheck(player *gameObject, enemies []*gameObject) {
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
			break
		}
	}
}

func drawEnemy(ctx js.Value, images map[string]js.Value, enemies []*gameObject) {
	for _, e := range enemies {
		ctx.Call("drawImage", images["player"], e.x-12, e.y-12)
	}
}

func drawPlayer(ctx js.Value, images map[string]js.Value, player *gameObject) {
	if player.alive {
		if player.frame%3 == 1 {
			// blink
			return
		}
		ctx.Call("drawImage", images["player"], player.x-12, player.y-12)
		return
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
