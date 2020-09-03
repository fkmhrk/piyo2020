package game

import (
	"syscall/js"
)

// Engine is an engine
type Engine interface {
	AddImage(key string, image js.Value)
	AddEnemy(enemy *gameObject)
	Player() *gameObject
	DoFrame(key int16, ctx js.Value)
}

type engine struct {
	images        map[string]js.Value
	player        *gameObject
	enemies       []*gameObject
	hiddenEnemies []*gameObject
	stage         *gameObject
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

func (e *engine) DoFrame(key int16, ctx js.Value) {
	movePlayer(e.player, key)
	moveEnemy(e.enemies, e)
	moveEnemy(e.hiddenEnemies, e)
	e.stage.moveFunc(e.stage, e)

	ctx.Call("clearRect", 0, 0, 320, 480)
	drawEnemy(ctx, e.images, e.enemies)
	ctx.Call("drawImage", e.images["player"], e.player.x, e.player.y)

	e.enemies = pack(e.enemies)
	e.hiddenEnemies = pack(e.hiddenEnemies)
}

func movePlayer(player *gameObject, key int16) {
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

func drawEnemy(ctx js.Value, images map[string]js.Value, enemies []*gameObject) {
	for _, e := range enemies {
		ctx.Call("drawImage", images["player"], e.x, e.y)
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
