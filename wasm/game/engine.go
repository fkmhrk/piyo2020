package game

import "syscall/js"

// Engine is an engine
type Engine interface {
	AddImage(key string, image js.Value)
	DoFrame(key int16, ctx js.Value)
}

type engine struct {
	images map[string]js.Value
	x      int
	y      int
}

// New creates engine instance
func New() Engine {
	return &engine{
		images: make(map[string]js.Value),
		x:      160,
		y:      440,
	}
}

func (e *engine) AddImage(key string, image js.Value) {
	e.images[key] = image
}

func (e *engine) DoFrame(key int16, ctx js.Value) {
	if key&1 != 0 {
		e.y -= 4
	}
	if key&2 != 0 {
		e.y += 4
	}
	if key&4 != 0 {
		e.x -= 4
	}
	if key&8 != 0 {
		e.x += 4
	}
	ctx.Call("clearRect", 0, 0, 320, 480)
	ctx.Call("drawImage", e.images["player"], e.x, e.y)
}
