package game

type Dead interface {
	SoloExplode() DeadFunc
	SoloExplodeWithItem(itemId int) DeadFunc
	Explode() DeadFunc
}
