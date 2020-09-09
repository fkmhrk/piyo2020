package game

type Dead interface {
	SoloExplode() DeadFunc
	SoloExplodeWithItem(itemId int) DeadFunc
	SoloExplodeWithItem3(itemId int) DeadFunc
	Explode() DeadFunc
}
