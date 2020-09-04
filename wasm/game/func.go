package game

import (
	"syscall/js"
)

type MoveFunc func(obj *GameObject, engine Engine)
type SeqMoveFunc func(obj *GameObject, engine Engine, frame int)

type SeqMove struct {
	Frame int
	Func  SeqMoveFunc
}

type SeqMoveFuncs []*SeqMove

type DeadFunc func(engine Engine, obj *GameObject)
type DrawFunc func(ctx js.Value, obj *GameObject, images map[string]*JsImage)

type ShotFunc func(obj *GameObject, engine Engine)
type SeqShotFunc func(obj *GameObject, engine Engine, frame int)

type SeqShot struct {
	Frame int
	Func  SeqShotFunc
}

type SeqShotFuncs []*SeqShot
