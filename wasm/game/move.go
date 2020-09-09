package game

type Move interface {
	Sequential() MoveFunc
	Nop() SeqMoveFunc

	Line() MoveFunc
	LineReflect() MoveFunc
	FrameCountUp() MoveFunc
	ItemDrop() MoveFunc
}
