package game

import (
	"due-v2-example/game/app/entity"
)

type Mahjong struct {
	table *entity.Table
}

func NewMahjong(table *entity.Table) *Mahjong {
	return &Mahjong{
		table: table,
	}
}
