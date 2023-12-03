package entity_test

import (
	"due-v2-example/game/app/entity"
	"fmt"
	"testing"
)

var (
	roomMgr   = entity.NewRoomMgr(nil)
	playerMgr = entity.NewPlayerMgr()
)

func TestNewRoomMgr(t *testing.T) {
	fmt.Println(roomMgr.GetTable(1, 1))
}

func TestRoomMgr_QuickMatch(t *testing.T) {
	player, err := playerMgr.LoadPlayer(6)
	if err != nil {
		t.Fatalf("load player failed: %v", err)
	}

	err = roomMgr.QuickMatch(player)
	if err != nil {
		t.Fatalf("quick match failed: %v", err)
	}

	t.Logf("room id: %d", player.Room().ID())
	t.Logf("table id: %d", player.Table().ID())
	t.Logf("seat id: %d", player.Seat().ID())
}
