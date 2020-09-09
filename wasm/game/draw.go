package game

type Draw interface {
	Static() DrawFunc
	StrokeArc() DrawFunc
	ExpandingStrokeArc() DrawFunc
	FillArc() DrawFunc
	StageText(stage int) DrawFunc
	Player() DrawFunc
}
